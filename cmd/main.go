package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/kkkfasya/fcaches/internal/cache"
	"github.com/kkkfasya/fcaches/internal/proxy"
)

var exampleOrigin string = "http://example.com"

type ProxyServer struct {
	port   int
	origin string
}

func main() {
	var ps ProxyServer

	flag.IntVar(&ps.port, "port", 6767, "Define port for proxy cache server to run on")
	flag.StringVar(&ps.origin, "origin", exampleOrigin, "Define origin URL of which proxy server will forward to")
	// CLEAR_CACHE := flag.Bool("clear-cache", false, "Clear current cache")

	flag.Parse()
	origin, err := url.Parse(ps.origin)
	if err != nil {
		log.Fatal(err)
	}
	proxy.NewProxy(origin)

}
