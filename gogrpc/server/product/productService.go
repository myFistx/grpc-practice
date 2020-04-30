package product

import (
	"context"
	product_proto "practice/gogrpc/server/proto_file"
)

type ProdService struct {
}

func (p *ProdService) GetProduct(ctx context.Context, req *product_proto.GetProductRequest) (*product_proto.GetProductResponse, error) {

	return &product_proto.GetProductResponse{
		ProdPrice: 20,
	}, nil
}
