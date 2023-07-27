package agent

import (
	"github.com/Rus-Iva/metriccollector/internal/dto"
	"github.com/Rus-Iva/metriccollector/internal/storage"
	"github.com/go-resty/resty/v2"
)

func (c *HTTPClient) sendMetricHandler(metricType, metricName, metricValue string) (*resty.Response, error) {
	resp, err := c.R().SetPathParams(map[string]string{
		"metricType":  metricType,
		"metricName":  metricName,
		"metricValue": metricValue,
	}).Post("/update/{metricType}/{metricName}/{metricValue}")
	return resp, err
}

func (c *HTTPClient) sendMetricJSONHandler(metricType string, metricName string, metricValue interface{}) (*resty.Response, error) {
	m := dto.Metrics{
		ID:    metricName,
		MType: metricType,
	}
	if v, ok := metricValue.(storage.Gauge); ok {
		m.Value = &v
	}
	if v, ok := metricValue.(storage.Counter); ok {
		m.Delta = &v
	}
	resp, err := c.R().SetHeader("Content-Type", "application/json").SetBody(m).Post("/update/")
	return resp, err
}
