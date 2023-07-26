package storage

type Storage interface {
	GetGauge() GaugeMetrics
	WriteGaugeValue(metricName string, metricVal float64)
	ReadGaugeValue(metricName string) (float64, error)

	GetCounter() CounterMetrics
	IncrementCounterValue(metricName string)
	WriteCounterValue(metricName string, metricVal int64) int64
	ReadCounterValue(metricName string) (int64, error)
}
