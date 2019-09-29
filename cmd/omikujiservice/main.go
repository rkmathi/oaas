package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"

	"oaas/internal/consts"
	pb "oaas/proto"
)

type omikuji struct{}

func (o *omikuji) DrawOmikuji(ctx context.Context, in *pb.DrawOmikujiRequest) (
	*pb.DrawOmikujiResponse, error) {
	log.Printf("START DrawOmikuji: %v", in)

	conn, err := grpc.Dial(consts.BankServicecAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}
	defer conn.Close()

	cli := pb.NewBankClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), consts.GrpcTimeout)
	defer cancel()

	res, err := cli.UpdateBalance(
		ctx, &pb.UpdateBalanceRequest{Delta: -consts.OmikujiCost})
	if err != nil {
		log.Fatalf("cannot UpdateBalance: %v", err)
	}

	switch res.Status.Code {
	case pb.Code_OK:
		t, _ := ptypes.Timestamp(in.Ts)
		status := &pb.Status{Code: pb.Code_OK, Message: "OK"}
		omikuji := &pb.OmikujiMessage{Value: drawOmikuji(in.Seed, t)}
		return &pb.DrawOmikujiResponse{Status: status, Omikuji: omikuji}, nil
	case pb.Code_BALANCE_INSUFFICIENT:
		status := &pb.Status{
			Code:    pb.Code_BALANCE_INSUFFICIENT,
			Message: pb.Code_name[int32(pb.Code_BALANCE_INSUFFICIENT)],
		}
		return &pb.DrawOmikujiResponse{Status: status}, nil
	}

	status := &pb.Status{
		Code:    pb.Code_UNKNOWN,
		Message: pb.Code_name[int32(pb.Code_UNKNOWN)],
	}
	return &pb.DrawOmikujiResponse{Status: status}, nil
}

func main() {
	log.Print("START omikuji service")

	listen, err := net.Listen("tcp", consts.OmikujiServicecAddr)
	if err != nil {
		log.Fatalf("cannot listen: %v", err)
	}

	svr := grpc.NewServer()
	pb.RegisterOmikujiServer(svr, &omikuji{})
	if err := svr.Serve(listen); err != nil {
		log.Fatalf("cannot serve: %v", err)
	}
}

func drawOmikuji(seed int32, t time.Time) string {
	if int(t.Month()) == 1 && t.Day() == 1 {
		return "dai-kichi"
	}

	// TODO: clever algorithm
	rand.Seed(int64(seed) + int64(t.Nanosecond()))
	switch rand.Intn(10) {
	case 0:
		return "dai-kichi"
	case 1, 2:
		return "chu-kichi"
	case 3, 4, 5, 6, 7:
		return "sho-kichi"
	case 8, 9:
		return "kyo"
	}

	return "???"
}
