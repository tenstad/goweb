package main

import (
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	httpRequestsProcessed := promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_processed",
		Help: "The total number of processed http requests",
	})

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		for i := 0; i < 10000; i++ {
			math.Pow(math.Pi, 1000)
		}
		end := time.Now()
		fmt.Println(fmt.Sprintf("%s %s", end.Format("2006-01-02 15:04:05.000000"), end.Sub(start)))

		fmt.Fprintf(w, "ok")
		httpRequestsProcessed.Inc()
	})

	http.ListenAndServe(":8080", nil)
}
