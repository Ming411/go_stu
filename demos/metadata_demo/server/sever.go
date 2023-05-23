package main

import (
	"context"
	"fmt"
	"go_stu/demos/metadata_demo/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	// 处理函数并返回
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println("get metadata error")
	}
	// for key, val := range md {
	// 	fmt.Println(key, val)
	// }
	// 如果name存在输出
	if nameSlice, ok := md["name"]; ok {
		fmt.Println(nameSlice) // nameSlice是一个切片
		for i, e := range nameSlice {
			fmt.Println(i, e)
		}
	}
	return &proto.HelloReply{
		Message: "Hello" + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer() // 开启grpc服务
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
