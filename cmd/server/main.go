package main

import (
	"net/http"
)

const (
	gauge   MetricType = "gauge"
	counter MetricType = "counter"
)

func main() {
	//ms := NewMemStorage()
	mux := http.NewServeMux()
	mux.HandleFunc(`/update/`, catchMetricHandler)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
