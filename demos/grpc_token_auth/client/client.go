package main

import (
	"context"
	"fmt"
	"go_stu/demos/grpc_token_auth/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type customCredential struct {
}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "01011",
		"appkey": "token",
	}, nil
}
func (c customCredential) RequireTransportSecurity() bool {
	return false
}

func main() {
	/* 方式一 */
	// interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 	start := time.Now()
	// 	md := metadata.New(map[string]string{
	// 		"appid":  "0101",
	// 		"appkey": "token",
	// 	})
	// 	ctx = metadata.NewOutgoingContext(context.Background(), md)
	// 	err := invoker(ctx, method, req, reply, cc, opts...)
	// 	fmt.Printf("消耗时间%s\n", time.Since(start))
	// 	return err
	// }
	// opt := grpc.WithUnaryInterceptor(interceptor)
	/* 方式二 */
	opt := grpc.WithPerRPCCredentials(customCredential{})
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
