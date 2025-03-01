package main

import (
	"fmt"
	"net/http"
	"net/url"

	xginx "github.com/xSaCh/xginx/pkg"
)

const PORT = 5000

func main() {
	lb := xginx.NewLoadBalancer()
	lb.AddBackend(&url.URL{Scheme: "http", Host: "localhost:8080"})
	lb.AddBackend(&url.URL{Scheme: "http", Host: "localhost:8081"})

	fmt.Printf("Running xginx on %d\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), http.HandlerFunc(lb.Router))
}
