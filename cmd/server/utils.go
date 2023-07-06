package main

import (
	"strconv"
	"strings"
)

func GetMetricParams(path string) (interface{}, error) {
	trimmedPath := strings.TrimPrefix(path, "/update/")
	pathParams := strings.Split(trimmedPath, "/")

	if len(pathParams) != 3 {
		return nil, &metricNameNotFoundError{path: path}
	}

	commonParams := commonMetricParams{
		metricType: MetricType(pathParams[0]),
		name:       pathParams[1],
	}

	if commonParams.metricType == gauge {
		metricVal, err := strconv.ParseFloat(pathParams[2], 64)
		if err != nil {
			return nil, &metricIncorrectTypeOrValue{path: path}
		}

		metricParams := gaugeMetricParams{
			commonMetricParams: commonParams,
			value:              metricVal,
		}
		return &metricParams, nil
	}
	if commonParams.metricType == counter {
		metricVal, err := strconv.ParseInt(pathParams[2], 2, 64)
		if err != nil {
			return nil, &metricIncorrectTypeOrValue{path: path}
		}

		metricParams := counterMetricParams{
			commonMetricParams: commonParams,
			value:              metricVal,
		}
		return &metricParams, nil
	}
	return nil, &metricIncorrectTypeOrValue{path: path}
}
