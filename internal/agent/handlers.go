package agent

import "github.com/go-resty/resty/v2"

func (c *HTTPClient) sendMetricHandler(metricType, metricName, metricValue string) (*resty.Response, error) {
	resp, err := c.R().SetPathParams(map[string]string{
		"metricType":  metricType,
		"metricName":  metricName,
		"metricValue": metricValue,
	}).Post("/update/{metricType}/{metricName}/{metricValue}")
	return resp, err
}
