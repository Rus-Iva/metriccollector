package main

import (
	"github.com/Rus-Iva/metriccollector/internal/agent"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	parseFlags()

	c := agent.NewClient(flagBseEndpointAddr)

	var shutdownCh = make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		c.CollectMetrics(time.Duration(flagPollInterval), time.Duration(flagReportInterval))
	}()

	<-shutdownCh
}
