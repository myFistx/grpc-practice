package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"practice/gogrpc/server/product"
	product_proto "practice/gogrpc/server/proto_file"
)

func main() {
	fmt.Println("grpc server start--")
	rpcServer := grpc.NewServer()
	product_proto.RegisterProductServiceServer(rpcServer, new(product.ProdService))

	listener, _ := net.Listen("tcp", ":8022")

	if err := rpcServer.Serve(listener); err != nil {
		fmt.Println(err)
	}
	fmt.Println(listener)
	fmt.Println("grpc server startup!")
}
