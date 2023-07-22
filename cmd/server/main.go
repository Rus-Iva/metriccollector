package main

import (
	"github.com/Rus-Iva/metriccollector/internal/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	s := server.NewServer()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/", func(r chi.Router) {
		r.Get("/", s.GetAllMetricsHandler)
		r.Post("/update/{metricType}/{metricName}/{metricValue}", s.PostMetricHandler)
		r.Get("/value/{metricType}/{metricName}", s.GetMetricValueHandler)
	})

	err := http.ListenAndServe(`:8080`, r)
	if err != nil {
		panic(err)
	}
}
