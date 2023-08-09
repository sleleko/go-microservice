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
	// for use via local machine
	//services := []Service{
	//	{Name: "service1", URL: "http://localhost:8081"},
	//	{Name: "service2", URL: "http://localhost:8082"},
	//}

	// for use via Docker
	services := []Service{
		{Name: "service1", URL: "http://service1:8081"},
		{Name: "service2", URL: "http://service2:8082"},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set("Content-Type", "text/html")
			// Отображаем страницу приветствия
			w.Write([]byte("This go microservice gateway, please select one of service to open:<br>"))
			for _, service := range services {
				w.Write([]byte(fmt.Sprintf("<a href='/%s' target='_blank'>%s</a><br>", service.Name, service.Name)))
			}
			return
		}

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
	parsedURL, _ := url.Parse(target) // Изменил имя переменной на parsedURL

	proxy := httputil.NewSingleHostReverseProxy(parsedURL) // Используем parsedURL

	r.URL.Host = parsedURL.Host
	r.URL.Scheme = parsedURL.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = parsedURL.Host

	proxy.ServeHTTP(w, r)
}
