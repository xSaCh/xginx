package pkg

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

type Backend struct {
	URL          *url.URL
	reverseProxy *httputil.ReverseProxy
	alive        bool
	lock         sync.RWMutex
}

func NewBackend(url *url.URL) *Backend {
	b := &Backend{
		URL:   url,
		alive: true,
	}
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

	b.reverseProxy = proxy
	b.alive = b.CheckAlive()
	return b
}

// Check if Backend is alive by dial tcp connection
func (b *Backend) CheckAlive() bool {
	conn, err := net.DialTimeout("tcp", b.URL.Host, time.Second*2)
	if err != nil {
		return false
	} else {
		defer conn.Close()
		return true
	}
}

func (b *Backend) IsAlive() bool {
	b.lock.RLock()
	alive := b.alive
	b.lock.RUnlock()
	return alive
}

func (b *Backend) SetAlive(alive bool) {
	b.lock.Lock()
	b.alive = alive
	b.lock.Unlock()
}

func (b *Backend) Serve(w http.ResponseWriter, r *http.Request) {
	b.reverseProxy.ServeHTTP(w, r)
}
