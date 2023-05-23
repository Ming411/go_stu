package main

import (
	"context"
	"fmt"
	"go_stu/demos/grpc_token_auth/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	// 处理函数并返回
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("get metadata error")
	}
	// for key, val := range md {
	// 	fmt.Println(key, val)
	// }
	// 如果name存在输出
	if nameSlice, ok := md["name"]; ok {
		// fmt.Println(nameSlice) // nameSlice是一个切片
		for i, e := range nameSlice {
			fmt.Println(i, e)
		}
	}
	return &proto.HelloReply{
		Message: "Hello" + request.Name,
	}, nil
}

func main() {
	// 拦截器
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("收到请求")
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			// status.Error()  grpc内部处理错误
			return resp, status.Error(codes.Unauthenticated, "无token认证信息")
		}
		var (
			appid  string
			appkey string
		)
		if va1, ok := md["appid"]; ok {
			appid = va1[0]
		}
		if va2, ok := md["appkey"]; ok {
			appkey = va2[0]
		}
		if appid != "0101" || appkey != "token" {
			return resp, status.Error(codes.Unauthenticated, "无token认证信息")
		}
		res, err := handler(ctx, req) //  原样返回之前的逻辑
		fmt.Println("请求完成")
		return res, err
	}
	opt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opt) // 开启grpc服务
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic("开启端口监听失败" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
