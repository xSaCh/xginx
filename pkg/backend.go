package xginx

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type Backend struct {
	URL          *url.URL
	Alive        bool
	ReverseProxy *httputil.ReverseProxy
	Lock         sync.RWMutex
}

func NewBackend(url *url.URL) *Backend {
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Error in [%s] for %s: %s", url, r.RequestURI, err.Error())
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
	}
	proxy.ModifyResponse = func(r *http.Response) error {
		lg := fmt.Sprintf("Response from %s for %s %s %s \n", url, r.Request.URL.RequestURI(), r.Proto, r.Status)
		log.Print(lg)
		return nil
	}
	return &Backend{
		URL:          url,
		Alive:        true,
		ReverseProxy: proxy,
	}
}

func (b *Backend) IsAlive() bool {
	return b.Alive
}

func (b *Backend) Serve(w http.ResponseWriter, r *http.Request) {
	b.ReverseProxy.ServeHTTP(w, r)
}
