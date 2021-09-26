package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	buckets            = []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10}
	numberOfRegistered = promauto.NewCounter(prometheus.CounterOpts{
		Name: "number_of_registered_users",
		Help: "The total number of registered users",
	})
	numberOfCakesGiven = promauto.NewCounter(prometheus.CounterOpts{
		Name: "number_of_cakes_given",
		Help: "The total number of cakes given",
	})

	responseTimeHistogram = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_server_request_duration_seconds",
		Help:    "Histogram of response time for handler in seconds",
		Buckets: buckets,
	}, []string{"route"})
)

func metrics() {

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)

	log.Println("If you see me, prometheus didn`t start")
}
