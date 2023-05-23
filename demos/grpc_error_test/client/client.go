package main

import (
	"context"
	"fmt"
	"go_stu/demos/grpc_error_test/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()                // 关闭连接
	c := proto.NewGreeterClient(conn) // 创建gRPC客户端
	// 设置超时时间
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	r, err := c.SayHello(ctx, &proto.HelloRequest{
		Name: "coder",
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			panic("解析error失败")
		}
		fmt.Println(st.Message())
		fmt.Println(st.Code())
	}
	fmt.Println(r.Message)
}
