package cache

import (
	"net/http"
	"time"
)

//i should be able to model the type like: type XCacheState = "HIT" | "MISS"
// but since golang type system is god awful...anyway sorry i'll stop
type XCacheState string
const (
	XCacheHit  XCacheState = "HIT"
	XCacheMiss XCacheState = "MISS"
)

type CacheObject struct {
	Response *http.Response
	Body     []byte
	Cached   time.Time
}

func NewCacheKey(r *http.Request) string {
	return r.Method + ":" + r.URL.String()
}

type MapCache map[string]*CacheObject
