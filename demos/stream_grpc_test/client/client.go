package mian

import (
	"context"
	"go_stu/demos/grpc_stu/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "immoc"})
}
