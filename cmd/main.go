package main

import (
	"context"
	"time"

	"github.com/k0kubun/pp"
	micro "github.com/micro/go-micro/v2"
	pbOrder "github.com/phanletrunghieu/demo-go-micro-order-srv/proto/order"
	pbPayment "github.com/phanletrunghieu/demo-go-micro-payment-srv/proto/payment"
)

func main() {
	var err error
	srv := micro.NewService(
		micro.Name("srv.cli"),
	)
	ctx := context.Background()

	orderClient := pbOrder.NewOrderService("srv.grpc.order", srv.Client())
	paymentClient := pbPayment.NewPaymentService("srv.grpc.payment", srv.Client())

	order := &pbOrder.Order{
		Id:        "111",
		CreatedAt: time.Now().UnixNano(),
	}
	o, err := orderClient.CreateOrder(ctx, order)
	if err != nil {
		panic(err)
	}
	pp.Println(o)

	transaction := &pbPayment.Transaction{
		Amount:    1000,
		CreatedAt: time.Now().UnixNano(),
		Id:        "111",
		OrderId:   "111",
	}
	t, err := paymentClient.CreateTransaction(ctx, transaction)
	if err != nil {
		panic(err)
	}
	pp.Println(t)
}
