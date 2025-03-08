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

	lb := xginx.NewLoadBalancer(&config.LoadBalancer)
	if lb == nil {
		log.Fatalf("can't create lb")
		return
	}
	if lb == nil {
		log.Fatalf("Couldn't create LB")
		return
	}
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

func main2() {
	config, err := xginx.LoadConfig("example_config.yaml")
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	fmt.Printf("Loaded config: %+v\n", config)

}
