package main

import (
	"github.com/Rus-Iva/metriccollector/internal/agent"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := agent.NewClient()

	var shutdownCh = make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		c.CollectMetrics()
	}()
	<-shutdownCh
}
