package server

import (
	"encoding/json"
	"fmt"
	"github.com/Rus-Iva/metriccollector/internal/models"
	"github.com/Rus-Iva/metriccollector/internal/storage"
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

func (s *Server) PostJSONMetricHandler(rw http.ResponseWriter, r *http.Request) {
	var metric models.Metrics
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&metric); err != nil {
		s.Logger.Error().Msg("cannot decode request JSON body")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respMetric := models.Metrics{
		ID:    metric.ID,
		MType: metric.MType,
	}
	if metric.MType == "gauge" {
		s.storage.WriteGaugeValue(metric.ID, *metric.Value)
		respMetric.Value = metric.Value
	} else if metric.MType == "counter" {
		actualVal := s.storage.WriteCounterValue(metric.ID, *metric.Delta)
		respMetric.Delta = &actualVal
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(rw)
	if err := enc.Encode(respMetric); err != nil {
		s.Logger.Error().Msg("error encoding response")
		return
	}
}

func (s *Server) PostMetricHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(rw, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}
	metricType := chi.URLParam(r, "metricType")
	metricName := chi.URLParam(r, "metricName")
	metricVal := chi.URLParam(r, "metricValue")

	if metricName == "" {
		http.Error(rw, "missed metric name", http.StatusNotFound)
		return
	}

	if metricType == "gauge" {
		metricValParsed, err := strconv.ParseFloat(metricVal, 64)
		if err != nil {
			http.Error(rw, "incorrect type of metric value", http.StatusBadRequest)
			return
		}
		s.storage.WriteGaugeValue(metricName, metricValParsed)
		rw.WriteHeader(http.StatusOK)
		return

	}
	if metricType == "counter" {
		metricValParsed, err := strconv.ParseInt(metricVal, 0, 64)
		if err != nil {
			http.Error(rw, "incorrect type of metric value", http.StatusBadRequest)
			return
		}
		s.storage.WriteCounterValue(metricName, metricValParsed)
		rw.WriteHeader(http.StatusOK)
		return
	}

	http.Error(rw, "incorrect metric type", http.StatusBadRequest)

}

func (s *Server) PostJSONMetricValueHandler(rw http.ResponseWriter, r *http.Request) {
	var metric models.Metrics
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&metric); err != nil {
		s.Logger.Error().Msg("cannot decode request JSON body")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respMetric := models.Metrics{
		ID:    metric.ID,
		MType: metric.MType,
	}
	if metric.MType == "gauge" {
		gaugeVal, err := s.storage.ReadGaugeValue(metric.ID)
		s.Logger.Info().Msg(fmt.Sprintf("resp JSON body, %s, %.2f", metric.ID, gaugeVal))
		if err != nil {
			http.Error(rw, "", http.StatusNotFound)
			return
		}
		respMetric.Value = &gaugeVal
	} else if metric.MType == "counter" {
		counterVal, err := s.storage.ReadCounterValue(metric.ID)
		s.Logger.Info().Msg(fmt.Sprintf("resp JSON body, %s, %d", metric.ID, counterVal))
		if err != nil {
			http.Error(rw, "", http.StatusNotFound)
			return
		}
		respMetric.Delta = &counterVal

	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(rw)
	if err := enc.Encode(respMetric); err != nil {
		s.Logger.Error().Msg("error encoding response")
		return
	}

}

func (s *Server) GetMetricValueHandler(rw http.ResponseWriter, r *http.Request) {
	metricType := chi.URLParam(r, "metricType")
	metricName := chi.URLParam(r, "metricName")
	if metricType == "gauge" {
		val, err := s.storage.ReadGaugeValue(metricName)
		if err != nil {
			http.Error(rw, "", http.StatusNotFound)
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(storage.GaugeString(val)))
	}
	if metricType == "counter" {
		val, err := s.storage.ReadCounterValue(metricName)
		if err != nil {
			http.Error(rw, "", http.StatusNotFound)
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(storage.CounterString(val)))
	}
}

func (s *Server) GetAllMetricsHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")

	exPath := filepath.Dir(s.executablePath)
	t := template.Must(template.ParseFiles(filepath.Join(exPath, "static/index.html")))
	context := Context{}
	gaugeMetrics := s.storage.GetGauge()
	for k, v := range gaugeMetrics {
		context.Gauge = append(context.Gauge, Metrics{k, storage.GaugeString(v)})
	}
	counterMetrics := s.storage.GetCounter()
	for k, v := range counterMetrics {
		context.Counter = append(context.Counter, Metrics{k, storage.CounterString(v)})
	}
	if tErr := t.Execute(rw, context); tErr != nil {
		panic(tErr)
	}
}
