package httphandler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	"oaas/internal/consts"
	hrr "oaas/internal/httprequestresponse"
	pb "oaas/proto"
)

func HandleBalance(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getBalance(w, r)
		return
	case http.MethodPut:
		updateBalance(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func getBalance(w http.ResponseWriter, r *http.Request) {
	log.Printf("START getBalance: %v", r)

	conn, err := grpc.Dial(consts.BankServicecAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}
	defer conn.Close()

	cli := pb.NewBankClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), consts.GrpcTimeout)
	defer cancel()

	res, err := cli.GetBalance(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("cannot GetBalance: %v", err)
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	hres := &hrr.GetBalanceResponse{
		Amount: res.Balance.Amount,
	}
	if err := enc.Encode(hres); err != nil {
		log.Fatal(err)
	}

	_, _ = fmt.Fprintf(w, buf.String())

	log.Print("END getBalance")
}

func updateBalance(w http.ResponseWriter, r *http.Request) {
	log.Printf("START updateBalance: %v", r)

	var bufBody bytes.Buffer
	_, err := bufBody.ReadFrom(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(&bufBody)
	var ubr hrr.UpdateBalanceRequest
	if err := dec.Decode(&ubr); err != nil {
		log.Fatalf("cannot decode: %v", err)
	}
	defer r.Body.Close()

	conn, err := grpc.Dial(consts.BankServicecAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}
	defer conn.Close()

	cli := pb.NewBankClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), consts.GrpcTimeout)
	defer cancel()

	res, err := cli.UpdateBalance(
		ctx, &pb.UpdateBalanceRequest{Delta: ubr.Delta})
	if err != nil {
		log.Fatalf("cannot UpdateBalance: %v", err)
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	hres := &hrr.UpdateBalanceResponse{
		Status: res.Status.Message,
		Amount: res.Balance.Amount,
	}
	if err := enc.Encode(hres); err != nil {
		log.Fatal(err)
	}

	_, _ = fmt.Fprintf(w, buf.String())

	log.Print("END updateBalance")
}
