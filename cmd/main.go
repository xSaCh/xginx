package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/xSaCh/xginx"
)

const PORT = 5000

func main() {
	config, err := xginx.LoadConfig("example_config.yaml")
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}
	fmt.Printf("Loaded config: %+v\n", config)

	lbConfig := config.LoadBalancer

	lb := xginx.NewLoadBalancer(&lbConfig)
	if lb == nil {
		log.Fatalf("couldn't create LB")
		return
	}
	ticker := time.NewTicker(time.Second * time.Duration(lbConfig.HealthCheck.Interval))
	defer ticker.Stop()

	for _, svr := range lbConfig.BackendServers {
		lb.AddBackend(&url.URL{Scheme: "http", Host: fmt.Sprintf("%s:%d", svr.Address, svr.Port)})
	}

	go func() {
		for range ticker.C {
			lb.ServerPool.HealthCheck()
		}

	}()

	fmt.Printf("Running xginx on %s:%d\n", lbConfig.Host, lbConfig.Port)
	for _, server := range lb.ServerPool.Servers {
		fmt.Printf("Server: %s\n", server.URL.String())
	}
	http.ListenAndServe(fmt.Sprintf("%s:%d", lbConfig.Host, lbConfig.Port), http.HandlerFunc(lb.Router))
}
