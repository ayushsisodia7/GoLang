package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"test_mod/api"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type myOrderCreationServer struct {
	api.UnimplementedOrderCreationServer
}

var (
	orderIDCounter int64
	orderIDMutex   sync.Mutex
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

func (s myOrderCreationServer) Create(ctx context.Context, req *api.CreateOrder) (*api.OrderResponse, error) {
	if req.Amount < "100" {
		return nil, status.Errorf(codes.InvalidArgument, "amount must be at least 100")
	}
	// Process the order here (e.g., save to a database)

	orderIDMutex.Lock()
	defer orderIDMutex.Unlock()
	orderIDCounter++

	return &api.OrderResponse{
		OrderId: fmt.Sprintf("order_%s", base36Encode(orderIDCounter)), //"order_%012d", orderIDCounter //uuid.New().String() //fmt.Sprintf("order_%s", base36Encode(orderIDCounter)) //fmt.Sprintf("order_%012d", orderIDCounter)

		Entity:    "order",
		Amount:    req.Amount,
		Status:    "created",
		Currency:  req.Currency,
		Receipt:   req.Receipt,
		OfferId:   req.OfferId,
		Notes:     req.Notes,
		CreatedAt: time.Now().Unix(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	fmt.Println("hello")

	serverRegistrar := grpc.NewServer()
	service := &myOrderCreationServer{}

	api.RegisterOrderCreationServer(serverRegistrar, service)

	err = serverRegistrar.Serve(lis)

	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
	// client := razorpay.NewClient("rzp_test_OQdHxtuaO0xSXt", "12SgLhJy9l6pJKw1Xs5Fwix4")

	// data := map[string]interface{}{
	// 	"amount":   50000,
	// 	"currency": "INR",
	// 	"receipt":  "some_receipt_id",
	// }

	// body, err := client.Order.Create(data, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(body)

	// pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	// logFatal(err)
	// db, err := jet.Open("postgres", pgUrl)
	// logFatal(err)

	// err = db.Query("SELECT * FROM people").Rows(&people)
	// logFatal(err)

}
