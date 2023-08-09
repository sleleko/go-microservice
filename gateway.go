package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Service represents a single microservice endpoint
type Service struct {
	Name string
	URL  string
}

func main() {
	services := []Service{
		{Name: "service1", URL: "http://localhost:8081"},
		{Name: "service2", URL: "http://localhost:8082"},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _, service := range services {
			if r.URL.Path == "/"+service.Name {
				serveReverseProxy(service.URL, w, r)
				return
			}
		}
		http.NotFound(w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveReverseProxy(target string, w http.ResponseWriter, r *http.Request) {
	url, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(url)

	r.URL.Host = url.Host
	r.URL.Scheme = url.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = url.Host

	proxy.ServeHTTP(w, r)
}
