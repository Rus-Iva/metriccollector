package storage

type Storage interface {
	GetGauge() GaugeMetrics
	WriteGaugeValue(metricName string, metricVal Gauge)
	ReadGaugeValue(metricName string) (Gauge, error)

	GetCounter() CounterMetrics
	IncrementCounterValue(metricName string)
	WriteCounterValue(metricName string, metricVal Counter)
	ReadCounterValue(metricName string) (Counter, error)
}
