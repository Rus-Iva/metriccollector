package storage

import (
	"fmt"
	"sync"
)

type MemData struct {
	Gauge   GaugeMetrics
	Counter CounterMetrics
}

type MemStorage struct {
	sync.RWMutex
	data *MemData
}

func (ms *MemStorage) GetGauge() GaugeMetrics {
	return ms.data.Gauge
}

func (ms *MemStorage) SetGauge(gm GaugeMetrics) {
	ms.data.Gauge = gm
}

func (ms *MemStorage) WriteGaugeValue(metricName string, metricVal float64) {
	//ms.Lock()
	//defer ms.Unlock()
	ms.data.Gauge[metricName] = metricVal
}

func (ms *MemStorage) ReadGaugeValue(metricName string) (float64, error) {
	//ms.RLock()
	//defer ms.RUnlock()
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

func (ms *MemStorage) WriteCounterValue(metricName string, metricVal int64) int64 {
	//ms.Lock()
	//defer ms.Unlock()
	if currVal, ok := ms.data.Counter[metricName]; ok {
		newVal := currVal + metricVal
		ms.data.Counter[metricName] = newVal
		return newVal
	}
	ms.data.Counter[metricName] = metricVal
	return metricVal
}

func (ms *MemStorage) ReadCounterValue(metricName string) (int64, error) {
	//ms.RLock()
	//defer ms.RUnlock()
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
