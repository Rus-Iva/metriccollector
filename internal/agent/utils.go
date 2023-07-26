package agent

import (
	"fmt"
	"github.com/Rus-Iva/metriccollector/internal/storage"
	"math/rand"
	"runtime"
	"time"
)

func (c *HTTPClient) pollMetrics() {
	runtime.ReadMemStats(c.myMemStats)
	gm := storage.GaugeMetrics{
		"Alloc":         float64(c.myMemStats.Alloc),
		"BuckHashSys":   float64(c.myMemStats.BuckHashSys),
		"Frees":         float64(c.myMemStats.Frees),
		"GCCPUFraction": c.myMemStats.GCCPUFraction,
		"GCSys":         float64(c.myMemStats.GCSys),
		"HeapAlloc":     float64(c.myMemStats.HeapAlloc),
		"HeapIdle":      float64(c.myMemStats.HeapIdle),
		"HeapInuse":     float64(c.myMemStats.HeapInuse),
		"HeapObjects":   float64(c.myMemStats.HeapObjects),
		"HeapReleased":  float64(c.myMemStats.HeapReleased),
		"HeapSys":       float64(c.myMemStats.HeapSys),
		"LastGC":        float64(c.myMemStats.LastGC),
		"Lookups":       float64(c.myMemStats.Lookups),
		"MCacheInuse":   float64(c.myMemStats.MCacheInuse),
		"MCacheSys":     float64(c.myMemStats.MCacheSys),
		"MSpanInuse":    float64(c.myMemStats.MSpanInuse),
		"MSpanSys":      float64(c.myMemStats.MSpanSys),
		"Mallocs":       float64(c.myMemStats.Mallocs),
		"NextGC":        float64(c.myMemStats.NextGC),
		"NumForcedGC":   float64(c.myMemStats.NumForcedGC),
		"NumGC":         float64(c.myMemStats.NumGC),
		"OtherSys":      float64(c.myMemStats.OtherSys),
		"PauseTotalNs":  float64(c.myMemStats.PauseTotalNs),
		"StackInuse":    float64(c.myMemStats.StackInuse),
		"StackSys":      float64(c.myMemStats.StackSys),
		"Sys":           float64(c.myMemStats.Sys),
		"TotalAlloc":    float64(c.myMemStats.TotalAlloc),
		"RandomValue":   float64(rand.Int()),
	}
	c.storage.SetGauge(gm)
	c.storage.IncrementCounterValue("PollCount")
}

func (c *HTTPClient) sendMetrics() {
	for k, v := range c.storage.GetGauge() {
		//resp, err := c.sendMetricHandler("gauge", k, storage.GaugeString(v))
		resp, err := c.sendJSONMetricHandler("gauge", k, v)
		if err != nil {
			panic(err)
		}
		fmt.Printf("key: %s, value: %.2f, resp status %d\n", k, v, resp.StatusCode())
	}
	for k, v := range c.storage.GetCounter() {
		//resp, err := c.sendMetricHandler("counter", k, storage.CounterString(v))
		resp, err := c.sendJSONMetricHandler("counter", k, v)
		if err != nil {
			panic(err)
		}
		fmt.Printf("key: %s, value: %d, resp status %d\n", k, v, resp.StatusCode())
	}
}

func (c *HTTPClient) CollectMetrics(pollInterval, reportInterval time.Duration) {
	pollTicker := time.NewTicker(pollInterval)
	reportTicker := time.NewTicker(reportInterval)

	go func() {
		for {
			select {
			case <-pollTicker.C:
				c.pollMetrics()
			case <-reportTicker.C:
				c.sendMetrics()
			}
		}
	}()

}
