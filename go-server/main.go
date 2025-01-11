package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cpuUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_usage_percent",
		Help: "Current CPU usage in percent.",
	})
	memoryUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memory_usage_mb",
		Help: "Current memory usage in MB.",
	})
)

func init() {
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(memoryUsage)
}

func main() {
	go func() {
		for {
			// пример генерации метрик
			cpuUsage.Set(getCPUUsage())
			memoryUsage.Set(getMemoryUsage())
			time.Sleep(2 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getCPUUsage() float64 {
	// здесь можно реализовать логику получения реальной нагрузки CPU
	return 45.7 // пример значения
}

func getMemoryUsage() float64 {
	// здесь можно реализовать логику получения реального использования памяти
	return 512.0 // пример значения
}