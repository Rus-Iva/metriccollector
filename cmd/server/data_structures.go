package main

type MetricType string

type Metrics struct {
	gauge   map[string]float64
	counter map[string]int64
}
type commonMetricParams struct {
	metricType MetricType
	name       string
}

type gaugeMetricParams struct {
	commonMetricParams commonMetricParams
	value              float64
}

type counterMetricParams struct {
	commonMetricParams commonMetricParams
	value              int64
}

type MemStorage struct {
	metrics Metrics
}

func NewMemStorage() *MemStorage {
	m := Metrics{
		gauge:   map[string]float64{},
		counter: map[string]int64{},
	}
	ms := MemStorage{
		metrics: m,
	}
	return &ms
}
