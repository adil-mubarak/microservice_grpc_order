package main

import (
	"fmt"
	"log"
	"microservice_grpc_order/db"
	"microservice_grpc_order/pb/order"
	"microservice_grpc_order/services"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db, err := db.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	grpcServer := grpc.NewServer()
	orderService := &services.OrderService{
		DB: db,
	}
	order.RegisterOrderServiceServer(grpcServer, orderService)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen port 8080: %v", err)
	}
	fmt.Println("Server running on port :8080")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
