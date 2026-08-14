package proxy

import (
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/kkkfasya/fcaches/internal/cache"
)



type ProxyObject struct {
	mu     sync.RWMutex
	Origin string
	Cache  cache.MapCache
}

func NewProxy(origin *url.URL) (*ProxyObject, error) {
	return &ProxyObject{
		Origin: origin.String(),
		Cache:  make(cache.MapCache),
	}, nil
}


func RespondWithHeader

func (p *ProxyObject) ClearCache() {
	p.mu.Lock()
	p.Cache = make(cache.MapCache)
	p.mu.Unlock()
	log.Print("cache cleared successfully") // TODO: use slog
}


func (p *ProxyObject) ServeProxy(w http.ResponseWriter, r *http.Request) {
	CACHE_KEY := cache.NewCacheKey(r)
	p.mu.RLock()
	if c, ok := p.Cache[CACHE_KEY]; ok {
		p.mu.RUnlock()

	}

	// allowMethod list is for method that does not need extra process
	// e.g POST, PUT, etc. if we get those method we need to do something with the cache
	allowMethod := map[string]bool{ // O(1) baby
		"GET":  true,
		"HEAD": true,
	}
	if allowMethod[r.Method] {

	}
}
