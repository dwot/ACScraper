package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalf("Failed to get IP addresses: %s", err)
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

func main() {
	proxy := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.String(), "http://www.acinfinityserver.com/api/") {
			token := r.Header.Get("token")
			userAgent := r.Header.Get("User-Agent")

			if token != "" {
				fmt.Printf("Captured Token: %s\n", token)
			}
			if userAgent != "" {
				fmt.Printf("User-Agent: %s\n", userAgent)
			}
		}

		resp, err := http.DefaultTransport.RoundTrip(r)
		if err != nil {
			http.Error(w, "Error making request to target", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		copyHeader(w.Header(), resp.Header)
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	})

	ip := getLocalIP()
	fmt.Println("Starting proxy server on IP:", ip, "and port: 8080")
	fmt.Println("On your phone, set your proxy server to:", ip, "with port: 8080")
	log.Fatal(http.ListenAndServe(":8080", proxy))
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
