package main

import (
	"fmt"
	"log"
	"time"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SayHello(greeting_summary *prometheus.SummaryVec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer r.Body.Close()
		code := 500

		defer func() {
			httpDuration := time.Since(start)
			greeting_summary.WithLabelValues("duration").Observe(httpDuration.Seconds())
		}()

		code = http.StatusBadRequest
		if r.Method == "GET" {
			code = http.StatusOK
			vars := mux.Vars(r)
			name := vars["name"]
			if (name == "slow") {
				time.Sleep(10 * time.Second)
			}

			greet := fmt.Sprintf("Hello %s \n", name)
			w.Write([]byte(greet))
		} else {
			w.WriteHeader(code)
		}
	}
}

func main() {
	greeting_summary := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "greeting_seconds",
		Help: "Time taken for greeting request",
	}, []string{"greetings"})
	
	router := mux.NewRouter()
	router.Handle("/hello/{name}", SayHello(greeting_summary))
	router.Handle("/metrics", promhttp.Handler())
	prometheus.Register(greeting_summary)
	log.Fatal(http.ListenAndServe(":8080", router))
}
