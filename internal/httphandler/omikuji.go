package httphandler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"

	"oaas/internal/consts"
	hrr "oaas/internal/httprequestresponse"
	pb "oaas/proto"
)

func HandleOmikuji(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createOmikuji(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func createOmikuji(w http.ResponseWriter, r *http.Request) {
	log.Printf("START createOmikuji: %v", r)

	conn, err := grpc.Dial(consts.OmikujiServicecAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}
	defer conn.Close()

	cli := pb.NewOmikujiClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), consts.GrpcTimeout)
	defer cancel()

	res, err := cli.DrawOmikuji(ctx,
		&pb.DrawOmikujiRequest{Seed: 1, Ts: ptypes.TimestampNow()})
	if err != nil {
		log.Fatalf("cannot DrawOmikuji: %v", err)
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	var hres *hrr.CreateOmikujiResponse
	switch res.Status.Code {
	case pb.Code_OK:
		hres = &hrr.CreateOmikujiResponse{
			Status: res.Status.Message,
			Value:  res.Omikuji.Value,
		}
	case pb.Code_BALANCE_INSUFFICIENT:
		hres = &hrr.CreateOmikujiResponse{
			Status: res.Status.Message,
			Value:  "",
		}
	}
	if err := enc.Encode(hres); err != nil {
		log.Fatal(err)
	}

	_, _ = fmt.Fprintf(w, buf.String())

	log.Print("END createOmikuji")
}
