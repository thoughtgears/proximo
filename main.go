package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"golang.org/x/oauth2/google"
)

func main() {
	targetURL := flag.String("url", "", "URL of the target service")
	flag.Parse()

	if *targetURL == "" {
		log.Fatal("Missing required flag: -url")
	}

	target, err := url.Parse(*targetURL)
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	// Modify the request's Host header to the target host.
	director := proxy.Director
	proxy.Director = func(req *http.Request) {
		director(req)
		req.Host = target.Host

		// Get JWT token
		token, err := getAccessToken(context.Background())
		if err != nil {
			log.Printf("Error getting access token: %v", err)
			return
		}

		// Add the Authorization header to the forwarded request.
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	http.HandleFunc("/", loggingMiddleware(handler(proxy)))

	log.Println("Starting server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p.ServeHTTP(w, r)
	}
}

func getAccessToken(ctx context.Context) (string, error) {
	credentials, err := google.FindDefaultCredentials(ctx, "your-required-scopes-here")
	if err != nil {
		log.Fatalf("Failed to find default credentials: %v", err)
	}

	ts := credentials.TokenSource
	token, err := ts.Token()
	if err != nil {
		log.Fatalf("Failed to retrieve token: %v", err)
	}

	return token.Extra("id_token").(string), nil
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: Method=%s, URL=%s", r.Method, r.URL)
		next.ServeHTTP(w, r) // call original handler
	}
}
