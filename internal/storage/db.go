package storage

import (
	"fmt"
)

type MemData struct {
	Gauge   GaugeMetrics
	Counter CounterMetrics
}

type MemStorage struct {
	data *MemData
}

func (ms *MemStorage) GetGauge() GaugeMetrics {
	return ms.data.Gauge
}

func (ms *MemStorage) SetGauge(gm GaugeMetrics) {
	ms.data.Gauge = gm
}

func (ms *MemStorage) WriteGaugeValue(metricName string, metricVal Gauge) {
	ms.data.Gauge[metricName] = metricVal
}

func (ms *MemStorage) ReadGaugeValue(metricName string) (Gauge, error) {
	val, ok := ms.data.Gauge[metricName]
	if ok {
		return val, nil
	}
	return val, fmt.Errorf("%s: metric name doesnt exist", metricName)
}

func (ms *MemStorage) GetCounter() CounterMetrics {
	return ms.data.Counter
}

func (ms *MemStorage) IncrementCounterValue(metricName string) {
	ms.data.Counter[metricName] = ms.data.Counter[metricName] + 1
}

func (ms *MemStorage) WriteCounterValue(metricName string, metricVal Counter) {
	if currVal, ok := ms.data.Counter[metricName]; ok {
		newVal := currVal + metricVal
		ms.data.Counter[metricName] = newVal
		return
	}
	ms.data.Counter[metricName] = metricVal
}

func (ms *MemStorage) ReadCounterValue(metricName string) (Counter, error) {
	val, ok := ms.data.Counter[metricName]
	if ok {
		return val, nil
	}
	return val, fmt.Errorf("%s: metric name doesnt exist", metricName)
}

func NewMemData() *MemData {
	md := MemData{
		Gauge:   NewGaugeMetrics(),
		Counter: NewCounterMetrics(),
	}
	return &md
}

func NewMemStorage() *MemStorage {
	ms := MemStorage{
		data: NewMemData(),
	}
	return &ms
}
