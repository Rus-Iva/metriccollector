package agent

import (
	"fmt"
	customlog "github.com/Rus-Iva/metriccollector/internal/logger"
	"github.com/Rus-Iva/metriccollector/internal/storage"
	"github.com/go-resty/resty/v2"
	"runtime"
	"time"
)

type HTTPClient struct {
	*resty.Client
	storage    *storage.MemStorage
	myMemStats runtime.MemStats
	Logger     *customlog.Logger
}

func NewClient(baseEndpoint string) *HTTPClient {
	client := resty.New()
	client.SetBaseURL(fmt.Sprintf("http://%s", baseEndpoint))
	client.SetHeader("Content-Type", "text/plain")
	client.SetTimeout(1 * time.Second)
	s := storage.NewMemStorage()
	logger := customlog.NewLogger()

	var m runtime.MemStats
	return &HTTPClient{client, s, m, logger}
}
