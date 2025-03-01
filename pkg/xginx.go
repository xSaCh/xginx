package xginx

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type LoadBalancer struct {
	ServerPool ServerPool
}

func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{}
}

func (lb *LoadBalancer) AddBackend(url *url.URL) {
	lb.ServerPool.Servers = append(lb.ServerPool.Servers, NewBackend(url))
}

func (lb *LoadBalancer) Router(w http.ResponseWriter, r *http.Request) {
	lg := fmt.Sprintf("Received request from %s\n", r.RemoteAddr)
	lg += fmt.Sprintf("%s %s %s\n", r.Method, r.RequestURI, r.Proto)
	for key, values := range r.Header {
		for _, value := range values {
			lg += fmt.Sprintf("%s: %s\n", key, value)
		}
	}
	log.Print(lg)
	lb.ServerPool.Servers[0].Serve(w, r)
}
