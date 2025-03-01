package xginx

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type Backend interface {
	IsAlive() bool
	Serve(http.ResponseWriter, *http.Request)
}

type backend struct {
	URL          *url.URL
	Alive        bool
	ReverseProxy *httputil.ReverseProxy
	Lock         sync.RWMutex
}

func NewBackend(url *url.URL) Backend {
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Error in [%s] for %s: %s", url, r.RequestURI, err.Error())
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
	}
	return &backend{
		URL:          url,
		Alive:        true,
		ReverseProxy: proxy,
	}
}

func (b *backend) IsAlive() bool {
	return b.Alive
}

func (b *backend) Serve(w http.ResponseWriter, r *http.Request) {

	b.ReverseProxy.ServeHTTP(w, r)
}

type ServerPool struct {
	Servers []Backend
}
