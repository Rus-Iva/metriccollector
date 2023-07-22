package main

import "fmt"

type metricNameNotFoundError struct {
	path string
}

func (mn *metricNameNotFoundError) Error() string {
	return fmt.Sprintf("parse %v: metric name not found in path params", mn.path)
}

type metricIncorrectTypeOrValue struct {
	path string
}

func (mit *metricIncorrectTypeOrValue) Error() string {
	return fmt.Sprintf("parse %v: incorrect type or value in path params", mit.path)
}
