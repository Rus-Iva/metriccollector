package storage

import (
	"strconv"
)

type GaugeMetrics map[string]float64
type CounterMetrics map[string]int64

func GaugeString(g float64) string {
	return strconv.FormatFloat(g, 'f', -1, 64)
}

func CounterString(c int64) string {
	return strconv.FormatInt(c, 10)
}

func NewGaugeMetrics() GaugeMetrics {
	gm := GaugeMetrics{
		"Alloc":         0,
		"BuckHashSys":   0,
		"Frees":         0,
		"GCCPUFraction": 0,
		"GCSys":         0,
		"HeapAlloc":     0,
		"HeapIdle":      0,
		"HeapInuse":     0,
		"HeapObjects":   0,
		"HeapReleased":  0,
		"HeapSys":       0,
		"LastGC":        0,
		"Lookups":       0,
		"MCacheInuse":   0,
		"MCacheSys":     0,
		"MSpanInuse":    0,
		"MSpanSys":      0,
		"Mallocs":       0,
		"NextGC":        0,
		"NumForcedGC":   0,
		"NumGC":         0,
		"OtherSys":      0,
		"PauseTotalNs":  0,
		"StackInuse":    0,
		"StackSys":      0,
		"Sys":           0,
		"TotalAlloc":    0,
		"RandomValue":   0,
	}
	return gm
}

func NewCounterMetrics() CounterMetrics {
	cm := CounterMetrics{
		"PollCount": 0,
	}
	return cm
}
