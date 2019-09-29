package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	"oaas/internal/consts"
	pb "oaas/proto"
)

var (
	// TODO: individual amount
	myAmount int32 = 0
)

type bank struct{}

func (b *bank) GetBalance(ctx context.Context, in *empty.Empty) (
	*pb.GetBalanceResponse, error) {
	log.Printf("START GetBalance  %v", in)

	balance := &pb.BalanceMessage{Amount: myAmount}
	status := &pb.Status{Code: pb.Code_OK, Message: "OK"}
	return &pb.GetBalanceResponse{Status: status, Balance: balance}, nil
}

func (b *bank) UpdateBalance(ctx context.Context, in *pb.UpdateBalanceRequest) (
	*pb.UpdateBalanceResponse, error) {
	log.Printf("START UpdateBalance  %v", in)

	// check amount
	if (myAmount + in.Delta) < 0 {
		status := &pb.Status{
			Code:    pb.Code_BALANCE_INSUFFICIENT,
			Message: "",
		}
		balance := &pb.BalanceMessage{Amount: myAmount}
		return &pb.UpdateBalanceResponse{Status: status, Balance: balance}, nil
	}

	myAmount += in.Delta

	balance := &pb.BalanceMessage{Amount: myAmount}
	status := &pb.Status{Code: pb.Code_OK, Message: "OK"}
	return &pb.UpdateBalanceResponse{Status: status, Balance: balance}, nil
}

func main() {
	log.Print("START bank service")

	listen, err := net.Listen("tcp", consts.BankServicecAddr)
	if err != nil {
		log.Fatalf("cannot listen: %v", err)
	}

	svr := grpc.NewServer()
	pb.RegisterBankServer(svr, &bank{})
	if err := svr.Serve(listen); err != nil {
		log.Fatalf("cannot serve: %v", err)
	}
}
