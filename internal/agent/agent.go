package agent

import (
	"fmt"
	"github.com/Rus-Iva/metriccollector/internal/storage"
	"github.com/go-resty/resty/v2"
	"runtime"
	"time"
)

type HTTPClient struct {
	*resty.Client
	storage    *storage.MemStorage
	myMemStats *runtime.MemStats
}

func NewClient(baseEndpoint string) *HTTPClient {
	client := resty.New()
	client.SetBaseURL(fmt.Sprintf("http://%s", baseEndpoint))
	client.SetHeader("Content-Type", "text/plain")
	client.SetTimeout(1 * time.Second)
	s := storage.NewMemStorage()

	var m runtime.MemStats
	return &HTTPClient{client, s, &m}
}
