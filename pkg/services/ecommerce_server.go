package services

import (
	"context"
	"db"
	"fmt"
	. "generated"

	"go.uber.org/zap"
)

type EcommerceServiceServer struct {
	UnimplementedEcommerceServer
	PrismaClient *db.PrismaClient
	Logger       *zap.SugaredLogger
}

func NewEcommerceServiceServer() *EcommerceServiceServer {
	return &EcommerceServiceServer{}
}

func (s *EcommerceServiceServer) ListProducts(ctx context.Context, req *ListProductsRequest) (*ListProductsResponse, error) {
	// TODO: Implement ListProducts
	products, err := s.PrismaClient.Product.FindMany().Exec(ctx)
	if err != nil {
		return nil, err
	}
	var productList []*Product
	for _, product := range products {
		productList = append(productList, &Product{
			Id:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
		})
	}
	return &ListProductsResponse{
		Products: productList,
	}, nil
}

func (s *EcommerceServiceServer) GetProduct(ctx context.Context, req *GetProductRequest) (*GetProductResponse, error) {
	// TODO: Implement GetProduct
	product, err := s.PrismaClient.Product.FindUnique(
		db.Product.ID.Equals(req.Id)).Exec(ctx)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, fmt.Errorf("product not found")
	}

	return &GetProductResponse{
		Product: &Product{
			Id:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
		},
	}, nil
}

func (s *EcommerceServiceServer) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*CreateOrderResponse, error) {
	// TODO: Implement CreateOrder
	order, err := s.PrismaClient.Order.CreateOne(
		db.Order.Status.Set("Pending"),
		db.Order.User.Link(
			db.User.ID.Equals(req.UserId),
		),
		db.Order.Products.Link(
			db.Product.ID.In(req.ProductIds),
		),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &CreateOrderResponse{
		Order: &Order{
			Id:         order.ID,
			UserId:     order.UserID,
			Status:     order.Status,
			ProductIds: req.ProductIds,
			Customer: &Customer{
				Id:    order.UserID,
				Name:  order.User().Name,
				Email: order.User().Email,
			},
		},
	}, nil
}
