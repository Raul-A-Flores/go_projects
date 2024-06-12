package main

import (
	"common"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRCP_ADDR", "localhost:2000")
)

func main() {

	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to liste: %v", err)
	}

	defer l.Close()

	store := NewStore()
	service := NewService(store)
	NewGRPHandler(grpcServer, service)

	log.Println("GRPC Server started at ", grpcAddr)

	service.CreateOrder(context.Background())

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
