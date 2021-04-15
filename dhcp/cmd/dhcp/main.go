package main

import (
	"context"
	"log"
	// "github.com/davecgh/go-spew/spew"
	"encoding/json"
	"github.com/insomniacslk/dhcp/dhcpv4/nclient4"
	"net/http"
	"os"
)

type ErrorResponse struct {
	Error string
}

type LeaseResponse struct {
	Ack    bool
	IP     string
	Server string
}

func main() {
	http.HandleFunc("/metrics", handleMetrics)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func handleMetrics(w http.ResponseWriter, r *http.Request) {
	iface := os.Getenv("IFACE")

	client, err := nclient4.New(iface, nclient4.WithDebugLogger())
	if err != nil {
		errorResponse(w, err, 500)
	}

	lease, err := client.Request(context.TODO())
	if err != nil {
		errorResponse(w, err, 500)
	}

	lr := LeaseResponse{
		Ack: lease.ACK != nil,
	}

	if lr.Ack {
		lr.IP = lease.ACK.YourIPAddr.String()
		lr.Server = lease.ACK.ServerIPAddr.String()
	}

	sendResponse(w, lr, http.StatusOK)
}

func sendResponse(w http.ResponseWriter, body interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Fatalf("failed to JSON encode response: %v", err)
	}
}

func errorResponse(w http.ResponseWriter, err error, status int) {
	r := ErrorResponse{
		Error: err.Error(),
	}

	sendResponse(w, r, status)
}
