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

type cacheResponse struct {
	body       []byte
	cacheState cache.XCacheState
	cacheKey   string // e.g GET:https://example.com
}

func NewProxy(origin *url.URL) (*ProxyObject, error) {
	return &ProxyObject{
		Origin: origin.String(),
		Cache:  make(cache.MapCache),
	}, nil
}

func RespondWithHeaders(w http.ResponseWriter, r *http.Request, resp cacheResponse) {
	w.Header().Set("X-Cache", string(resp.cacheState))
	w.WriteHeader(r.Response.StatusCode)
	for k, v := range r.Response.Header {
		w.Header()[k] = v
	}
	w.Write(resp.body)
}

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
