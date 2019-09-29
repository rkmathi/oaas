package main

import (
	"log"
	"net/http"

	"oaas/internal/consts"
	handler "oaas/internal/httphandler"
)

func main() {
	log.Print("START gateway")

	http.HandleFunc("/balance", handler.HandleBalance)
	http.HandleFunc("/omikuji", handler.HandleOmikuji)

	err := http.ListenAndServe(consts.GatewayAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
