package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/xSaCh/xginx"
	"github.com/xSaCh/xginx/pkg/schedulers"
)

const PORT = 5000

func main() {
	config := &xginx.LoadBalancerConfig{
		Scheduler: schedulers.SCHEDULER_ROUND_ROBIN,
	}
	lb := xginx.NewLoadBalancer(config)
	ticker := time.NewTicker(time.Minute * 2)
	defer ticker.Stop()

	lb.AddBackend(&url.URL{Scheme: "http", Host: "localhost:8080"})
	lb.AddBackend(&url.URL{Scheme: "http", Host: "localhost:8081"})
	lb.AddBackend(&url.URL{Scheme: "http", Host: "localhost:8082"})

	go func() {
		for {
			select {
			case <-ticker.C:
				lb.ServerPool.HealthCheck()
			}
		}
	}()

	fmt.Printf("Running xginx on %d\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), http.HandlerFunc(lb.Router))
}
