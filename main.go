package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Main function
func main() {
	// Check args
	targetFuzzers, targetErr := getFuzzersToWatch()
	if targetErr != nil {
		log.Println(targetErr.Error())
		return
	}

	// Start thread to watch the fuzzer(s)
	registerFuzzers(targetFuzzers)
	go watchFuzzers()

	// Start HTTP handler exposing the metrics
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
