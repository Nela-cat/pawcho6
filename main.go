package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

var version = "unknown"

func handler(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	if iphost, _, err := net.SplitHostPort(ip); err == nil {
		ip = iphost
	}
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "IP: %s\nHostname: %s\nVersion: %s\n", ip, hostname, version)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
