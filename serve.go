package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	fs := http.FileSystem(http.Dir("app/build/outputs/apk/debug"))
	http.Handle("/", http.StripPrefix("/", http.FileServer(fs)))

	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			fmt.Printf("Serving APKs on http://%s:8088\n", ipnet.IP.String())
		}
	}
	http.ListenAndServe(":8088", nil)
}
