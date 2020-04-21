package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	requests := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "requests",
			Help: "A counter for the requests signal",
		},
		[]string{"code", "method"},
	)

	duration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "duration",
			Help:    "A histogram of latencies for the duration signal",
			Buckets: []float64{.025, .05, 0.1, 0.25, 0.5, 0.75},
		},
		[]string{"handler", "method"},
	)

	timedSignalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sleepTime := rand.Intn(1000)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		elapsed := time.Since(start)
		fmt.Fprintln(w, "Request Completed - duration %s", elapsed)
		log.Printf("Request Completed - duration %s", elapsed)
	})

	errorSignalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sleepTime := rand.Intn(1000)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		elapsed := time.Since(start)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Printf("Request Errored - duration %s", elapsed)
	})

	timedSignalChain := promhttp.InstrumentHandlerDuration(duration.MustCurryWith(prometheus.Labels{"handler": "signals"}),
		promhttp.InstrumentHandlerCounter(requests, timedSignalHandler),
	)

	errorSignalChain := promhttp.InstrumentHandlerDuration(duration.MustCurryWith(prometheus.Labels{"handler": "signals"}),
		promhttp.InstrumentHandlerCounter(requests, errorSignalHandler),
	)

	prometheus.MustRegister(requests, duration)

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/fast_response", timedSignalChain)
	http.Handle("/error_response", errorSignalChain)
	s := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":2112",
	}
	log.Printf("4 Golden Signals Server Listening")
	s.ListenAndServe()
}
