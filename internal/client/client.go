package client

import (
	"github.com/Rus-Iva/metriccollector/internal/storage"
	"github.com/go-resty/resty/v2"
	"runtime"
	"time"
)

const (
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second
)

type HTTPClient struct {
	*resty.Client
	storage    *storage.MemStorage
	myMemStats runtime.MemStats
}

func NewClient() *HTTPClient {
	client := resty.New()
	client.SetBaseURL("http://localhost:8080")
	client.SetHeader("Content-Type", "text/plain")
	client.SetTimeout(1 * time.Second)
	s := storage.NewMemStorage()

	var m runtime.MemStats
	return &HTTPClient{client, s, m}
}
