package main

import (
	"github.com/Rus-Iva/metriccollector/internal/agent"
	"github.com/go-resty/resty/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	parseFlags()

	c := agent.NewClient(flagBaseEndpointAddr)
	c.OnError(func(req *resty.Request, err error) {
		if v, ok := err.(*resty.ResponseError); ok {
			// Do something with v.Response
			c.Logger.Error().Str("RESPONSE", v.Response.String())
		}
	})

	var shutdownCh = make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		c.CollectMetrics(time.Duration(flagReportInterval)*time.Second, time.Duration(flagReportInterval)*time.Second)
	}()

	<-shutdownCh
}
