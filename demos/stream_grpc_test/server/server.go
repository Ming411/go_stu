package main

import (
	"fmt"
	"go_stu/demos/stream_grpc_test/proto"
	"net"
	"time"

	"google.golang.org/grpc"
)

const PORT = ":50052"

type server struct {
}

// 这里拥有的方法必须与 proto 内的保持一致
// 至于需要什么参数  查看proto生成的.go文件
func (s *server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = res.Send(&proto.StreamResData{
			// 返回当前时间戳
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}
func (s *server) PutStream(cliStr proto.Greeter_PutStreamServer) error {
	return nil
}
func (s *server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	return nil
}
func (*server) mustEmbedUnimplementedGreeterServer() {}
func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})

	_ = s.Serve((lis))
}
