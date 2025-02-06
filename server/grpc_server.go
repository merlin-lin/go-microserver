package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"go-microserver/server/proto"
)

// 实现 HelloGRPC 服务
type server struct{}

func (s *server) HelloGRPC(ctx context.Context, req *proto.HelloReq) (*proto.HelloResp, error) {
	return &proto.HelloResp{
		Message: "Hello, " + req.Name,
	}, nil
}

// 启动 gRPC 服务器
func StartGRPCServer(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// 注册服务
	proto.RegisterHelloGRPCServer(grpcServer, new(server))

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
