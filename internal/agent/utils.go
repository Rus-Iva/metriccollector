package agent

import (
	"github.com/Rus-Iva/metriccollector/internal/storage"
	"math/rand"
	"runtime"
	"time"
)

func (c *HTTPClient) pollMetrics() {
	runtime.ReadMemStats(&c.myMemStats)
	gm := storage.GaugeMetrics{
		"Alloc":         storage.Gauge(c.myMemStats.Alloc),
		"BuckHashSys":   storage.Gauge(c.myMemStats.BuckHashSys),
		"Frees":         storage.Gauge(c.myMemStats.Frees),
		"GCCPUFraction": storage.Gauge(c.myMemStats.GCCPUFraction),
		"GCSys":         storage.Gauge(c.myMemStats.GCSys),
		"HeapAlloc":     storage.Gauge(c.myMemStats.HeapAlloc),
		"HeapIdle":      storage.Gauge(c.myMemStats.HeapIdle),
		"HeapInuse":     storage.Gauge(c.myMemStats.HeapInuse),
		"HeapObjects":   storage.Gauge(c.myMemStats.HeapObjects),
		"HeapReleased":  storage.Gauge(c.myMemStats.HeapReleased),
		"HeapSys":       storage.Gauge(c.myMemStats.HeapSys),
		"LastGC":        storage.Gauge(c.myMemStats.LastGC),
		"Lookups":       storage.Gauge(c.myMemStats.Lookups),
		"MCacheInuse":   storage.Gauge(c.myMemStats.MCacheInuse),
		"MCacheSys":     storage.Gauge(c.myMemStats.MCacheSys),
		"MSpanInuse":    storage.Gauge(c.myMemStats.MSpanInuse),
		"MSpanSys":      storage.Gauge(c.myMemStats.MSpanSys),
		"Mallocs":       storage.Gauge(c.myMemStats.Mallocs),
		"NextGC":        storage.Gauge(c.myMemStats.NextGC),
		"NumForcedGC":   storage.Gauge(c.myMemStats.NumForcedGC),
		"NumGC":         storage.Gauge(c.myMemStats.NumGC),
		"OtherSys":      storage.Gauge(c.myMemStats.OtherSys),
		"PauseTotalNs":  storage.Gauge(c.myMemStats.PauseTotalNs),
		"StackInuse":    storage.Gauge(c.myMemStats.StackInuse),
		"StackSys":      storage.Gauge(c.myMemStats.StackSys),
		"Sys":           storage.Gauge(c.myMemStats.Sys),
		"TotalAlloc":    storage.Gauge(c.myMemStats.TotalAlloc),
		"RandomValue":   storage.Gauge(rand.Int()),
	}
	c.storage.SetGauge(gm)
	c.storage.IncrementCounterValue("PollCount")
}

func (c *HTTPClient) sendMetrics() {
	for k, v := range c.storage.GetGauge() {
		//resp, err := c.sendMetricHandler("gauge", k, v.String())
		resp, err := c.sendMetricJSONHandler("gauge", k, v)
		if err != nil {
			c.Logger.Error().Err(err)
		}
		c.Logger.Info().Str("key", k).Str("value", v.String()).Int("status_code", resp.StatusCode())
	}
	for k, v := range c.storage.GetCounter() {
		//resp, err := c.sendMetricHandler("counter", k, v.String())
		resp, err := c.sendMetricJSONHandler("counter", k, v)
		if err != nil {
			c.Logger.Error().Err(err)
		}
		c.Logger.Info().Str("key", k).Str("value", v.String()).Int("status_code", resp.StatusCode())
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
