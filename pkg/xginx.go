package xginx

import (
	"fmt"
	"log"
	"net/http"
)

type LoadBalancer struct {
}

func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{}
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
	w.Write(fmt.Appendf(nil, "Hello from %s", r.RequestURI))
}
