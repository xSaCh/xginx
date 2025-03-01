package xginx

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/xSaCh/xginx/pkg"
	"github.com/xSaCh/xginx/pkg/schedulers"
)

type LoadBalancerConfig struct {
	Scheduler schedulers.Schedulers
}

type LoadBalancer struct {
	ServerPool pkg.ServerPool
	Config     LoadBalancerConfig
	scheduler  schedulers.Scheduler
}

func NewLoadBalancer(config *LoadBalancerConfig) *LoadBalancer {
	lb := &LoadBalancer{
		ServerPool: pkg.ServerPool{},
		Config:     *config,
	}
	lb.scheduler = schedulers.NewScheduler(config.Scheduler, &lb.ServerPool)
	return lb
}

func (lb *LoadBalancer) AddBackend(url *url.URL) {
	lb.ServerPool.Servers = append(lb.ServerPool.Servers, pkg.NewBackend(url))
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

	b := lb.scheduler.GetNextBackend()
	if b == nil {
		fmt.Println("NILL")
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	b.Serve(w, r)
}

func (lb *LoadBalancer) TT() {
	q := lb.scheduler.T()
	if q == nil {
		fmt.Println("NULLL")
	}
	fmt.Printf("%d \n", len(q.Servers))
}
