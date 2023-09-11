package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"toll-calculator/types"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "the isten addres of the HTTP server.")
	flag.Parse()

	str := NewMemoryStore()
	svc := NewInvoiceAggregator(str)

	makeHTTPTransport(*listenAddr, svc)

	fmt.Println("working fine")
}

func makeHTTPTransport(listenAddr string, svc Aggregator) {
	fmt.Println("HTTP transport running on port ", listenAddr)
	http.HandleFunc("/aggregate", handleAggregate(svc))
	http.ListenAndServe(listenAddr, nil)
}

func handleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
