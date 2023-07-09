package main

import (
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
