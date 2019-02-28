package grpc

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"

	"github.com/hikingpig/microservice-simple-flow/pkg/api/v1"
)

func RunServer(ctx context.Context, v1API v1.ToDoServiceServer, port string) error {
	listen, _ := net.Listen("tcp", ":"+port)
	// v1API is the protobuf
	// we register it with the server so the message type and rpc service can pass
	// that is we just register an interface
	// the real service is not registered yet
	server := grpc.NewServer()
	v1.RegisterToDoServiceServer(server, v1API)

	log.Println("starting gRPC server...")
	return server.Serve(listen)
	//server.Server will start another goroutine
	//and return nil error
}
