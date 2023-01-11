package web

import (
	"log"
	"net/http"
	"time"
)

func logger(h http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	getIP := func(r *http.Request) string {
		remoteIP := r.RemoteAddr

		if forwarded := r.Header["X-Forwarded-For"]; len(forwarded) != 0 {
			remoteIP = forwarded[len(forwarded)-1]
		}

		return remoteIP
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ip := getIP(r)
		now := time.Now()

		log.Printf("Started %s %s %s\n", ip, r.Method, r.URL)
		h(w, r)
		log.Printf("Completed %s %s %s in %s\n", ip, r.Method, r.URL, time.Since(now))
	}
}
