package main

import (
	"github.com/Rus-Iva/metriccollector/internal/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

func main() {
	parseFlags()

	s := server.NewServer()

	r := chi.NewRouter()
	r.Use(s.LoggerMiddleware)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.Recoverer)

	r.Route("/", func(r chi.Router) {
		r.Get("/", s.GetAllMetricsHandler)
		r.Post("/update/{metricType}/{metricName}/{metricValue}", s.PostMetricHandler)
		r.Get("/value/{metricType}/{metricName}", s.GetMetricValueHandler)
		r.Post("/update/", s.PostUpdateJSONMetricHandler)
		r.Post("/value/", s.PostValueJSONMetricHandler)
	})
	err := http.ListenAndServe(flagRunAddr, r)
	if err != nil {
		panic(err)
	}
}
