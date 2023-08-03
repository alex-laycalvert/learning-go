package main

import (
	"log"
	"net/http"
)

func main() {
	go http.ListenAndServe(":80", http.HandlerFunc(redirect))
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/favicon.ico", index)
	log.Fatal(http.ListenAndServeTLS(":443", "ca.crt", "ca.key", mux))
}

func index(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		http.ServeFile(w, req, "index.html")
	case "/favicon.ico":
		http.ServeFile(w, req, "favicon.ico")
	default:
		log.Printf("404: %s", req.URL.String())
		http.NotFound(w, req)
	}
}

// NOTE: https://gist.github.com/d-schmidt/587ceec34ce1334a5e60
func redirect(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target,
		// see comments below and consider the codes 308, 302, or 301
		http.StatusTemporaryRedirect)
}
