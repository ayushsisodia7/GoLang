package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"retail/shop_code"
	"sync"

	"google.golang.org/grpc"
)

type myProductServer struct {
	shop_code.UnimplementedProductServer
}

type myOrderServer struct {
	shop_code.UnimplementedOrderServer
}

var (
	orderIDCounter    int64
	orderIDMutex      sync.Mutex
	productIDCounter  int64
	productIDMutex    sync.Mutex
	customerIDCounter int64
	customerIDMutex   sync.Mutex
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func base36Encode(num int64) string {
	chars := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, 14)
	for i := 13; i >= 0; i-- {
		result[i] = chars[num%36]
		num /= 36
	}
	return string(result)
}

func (s myProductServer) CreatingProduct(ctx context.Context, req *shop_code.CreateProduct) (res *shop_code.ProductResponse, err error) {

	productIDMutex.Lock()
	defer productIDMutex.Unlock()
	productIDCounter++

	return &shop_code.ProductResponse{
		ProductId:   fmt.Sprintf("prod_%s", base36Encode(productIDCounter)),
		ProductName: req.ProductName,
		Price:       req.Price,
		Quantity:    req.Quantity,
		Message:     "Product has been created successfully",
	}, nil
}

func (s myProductServer) FetchingProducts(context.Context, *shop_code.Blank) (*shop_code.Products, error) {
	return &shop_code.Products{
		Product: "Products to be displayed here",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	fmt.Println("hello")

	serverRegistrar := grpc.NewServer()
	service := &myProductServer{}

	shop_code.RegisterProductServer(serverRegistrar, service)

	err = serverRegistrar.Serve(lis)

	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
