package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	product_proto "practice/gogrpc/server/proto_file"
)

func main() {
	conn, err := grpc.Dial(":8022", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	prodClient := product_proto.NewProductServiceClient(conn)
	res, err2 := prodClient.GetProduct(context.Background(), &product_proto.GetProductRequest{
		ProdId: "12",
	})
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(res)
	}

}
