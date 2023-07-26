package agent

import (
	"github.com/Rus-Iva/metriccollector/internal/models"
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

func (c *HTTPClient) sendJSONMetricHandler(metricType string, metricName string, metricValue interface{}) (*resty.Response, error) {
	reqData := models.Metrics{
		ID:    metricName,
		MType: metricType,
	}
	if v, ok := metricValue.(float64); ok {
		reqData.Value = &v
	}
	if v, ok := metricValue.(int64); ok {
		reqData.Delta = &v
	}

	resp, err := c.R().SetHeader("Content-Type", "application/json").SetBody(&reqData).Post("/update/")
	return resp, err
}
