package main

import (
	"context"
	"fmt"
	"go_stu/demos/metadata_demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	// 建立 grpc 连接
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
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
