package server

import (
	"bytes"
	"encoding/json"
	"github.com/Rus-Iva/metriccollector/internal/dto"
	"github.com/Rus-Iva/metriccollector/internal/storage"
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

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
		s.storage.WriteGaugeValue(metricName, storage.Gauge(metricValParsed))
		rw.WriteHeader(http.StatusOK)
		return

	}
	if metricType == "counter" {
		metricValParsed, err := strconv.ParseInt(metricVal, 0, 64)
		if err != nil {
			http.Error(rw, "incorrect type of metric value", http.StatusBadRequest)
			return
		}
		s.storage.WriteCounterValue(metricName, storage.Counter(metricValParsed))
		rw.WriteHeader(http.StatusOK)
		return
	}

	http.Error(rw, "incorrect metric type", http.StatusBadRequest)

}

func (s *Server) PostUpdateJSONMetricHandler(rw http.ResponseWriter, r *http.Request) {
	var m dto.Metrics
	var buf bytes.Buffer
	// читаем тело запроса
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(buf.Bytes(), &m); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	respMetrics := dto.Metrics{
		ID:    m.ID,
		MType: m.MType,
	}
	if m.MType == "gauge" {
		s.storage.WriteGaugeValue(m.ID, *m.Value)
		respMetrics.Value = m.Value
	} else if m.MType == "counter" {
		updatedVal := s.storage.WriteCounterValue(m.ID, *m.Delta)
		respMetrics.Delta = &updatedVal
	}

	resp, err := json.Marshal(respMetrics)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(resp)
}

func (s *Server) PostValueJSONMetricHandler(rw http.ResponseWriter, r *http.Request) {
	var m dto.Metrics
	var buf bytes.Buffer
	// читаем тело запроса
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(buf.Bytes(), &m); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	respMetrics := dto.Metrics{
		ID:    m.ID,
		MType: m.MType,
	}
	if m.MType == "gauge" {
		val, err := s.storage.ReadGaugeValue(m.ID)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusNotFound)
			return
		}
		respMetrics.Value = &val
	} else if m.MType == "counter" {
		val, err := s.storage.ReadCounterValue(m.ID)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusNotFound)
			return
		}
		respMetrics.Delta = &val
	}

	resp, err := json.Marshal(respMetrics)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(resp)
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
		rw.Write([]byte(val.String()))
	}
	if metricType == "counter" {
		val, err := s.storage.ReadCounterValue(metricName)
		if err != nil {
			http.Error(rw, "", http.StatusNotFound)
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(val.String()))
	}
}

func (s *Server) GetAllMetricsHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")

	exPath := filepath.Dir(s.executablePath)
	t := template.Must(template.ParseFiles(filepath.Join(exPath, "static/index.html")))
	context := Context{}
	gaugeMetrics := s.storage.GetGauge()
	for k, v := range gaugeMetrics {
		context.Gauge = append(context.Gauge, Metrics{k, v.String()})
	}
	counterMetrics := s.storage.GetCounter()
	for k, v := range counterMetrics {
		context.Counter = append(context.Counter, Metrics{k, v.String()})
	}
	if tErr := t.Execute(rw, context); tErr != nil {
		panic(tErr)
	}
}
