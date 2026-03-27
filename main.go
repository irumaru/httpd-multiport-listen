package main

import (
	"log"
	"net"
	"net/http"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	// root endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		port := ""
		if addr, ok := r.Context().Value(http.LocalAddrContextKey).(net.Addr); ok {
			_, p, err := net.SplitHostPort(addr.String())
			if err == nil {
				port = p
			}
		}
		if port == "" {
			_, p, err := net.SplitHostPort(r.Host)
			if err == nil {
				port = p
			} else {
				port = r.Host
			}
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("port=" + port))
	})

	// health check endpoint
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	addrs := []string{":8080", ":8081"}
	for i, addr := range addrs {
		if i == 0 {
			log.Printf("listening on %s", addr)
			continue
		}

		go func(a string) {
			log.Printf("listening on %s", a)
			if err := http.ListenAndServe(a, nil); err != nil {
				log.Fatalf("server failed: %v", err)
			}
		}(addr)
	}

	if err := http.ListenAndServe(addrs[0], nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
