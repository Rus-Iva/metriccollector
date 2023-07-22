package main

import "net/http"

func catchMetricHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(rw, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	_, err := GetMetricParams(r.URL.Path)
	if err != nil {
		if _, ok := err.(*metricNameNotFoundError); ok {
			rw.WriteHeader(http.StatusNotFound)
			return
		} else if _, ok := err.(*metricIncorrectTypeOrValue); ok {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	rw.WriteHeader(http.StatusOK)
}
