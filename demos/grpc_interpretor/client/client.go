package main

import (
	"context"
	"fmt"
	"go_stu/demos/grpc_interpretor/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {

	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Printf("消耗时间%s\n", time.Since(start))
		return err
	}
	opt := grpc.WithUnaryInterceptor(interceptor)
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure(), opt)
	if err != nil {
		panic(err)
	}
	defer conn.Close()                // 关闭连接
	c := proto.NewGreeterClient(conn) // 创建gRPC客户端
	md := metadata.New(map[string]string{
		"name":    "coder",
		"pasword": "imooc",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.SayHello(ctx, &proto.HelloRequest{
		Name: "whyccc",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
