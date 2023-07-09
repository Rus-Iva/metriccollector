package storage

import "strconv"

type Gauge float64
type Counter int64

type GaugeMetrics map[string]Gauge
type CounterMetrics map[string]Counter

func (g Gauge) String() string {
	return strconv.FormatFloat(float64(g), 'f', -1, 64)
}

func (c Counter) String() string {
	return strconv.FormatInt(int64(c), 10)
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
