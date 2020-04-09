package main

import (
	"net/http"
)

import (
	"fmt"
	"github.com/domac/playflame/stats"
	"github.com/varstr/uaparser"
	"os"
	"path/filepath"
	"strings"
    "time"
	_ "net/http/pprof"
	"net/http"
)

func main() {


	go func() {
		log.Logger.Info("start pprof: %v", http.ListenAndServe(":8902", nil))
	}()

    //高级接口
    http.HandleFunc("/advance", WithAdvanced(Simple))

    //简单接口
    http.HandleFunc("/simple", Simple)
    // http.HandleFunc("/", index)

    hostPort := ":9090"
    fmt.Println("Starting Server on", hostPort)
    if err := http.ListenAndServe(hostPort, nil); err != nil {
        // log.Fatalf("HTTP Server Failed: %v", err)
        fmt.Println("ListenAndServe() error")
        return
    }
}


func Simple(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello VIP!")
}

//高级封装
func WithAdvanced(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		tags := getStatsTags(r)
		stats.IncCounter("handler.received", tags, 1)
		h(w, r)
		duration := time.Since(start)
		stats.RecordTimer("handler.latency", tags, duration)
	}
}

func getStatsTags(r *http.Request) map[string]string {
	userBrowser, userOS := parseUserAgent(r.UserAgent())
	stats := map[string]string{
		"browser":  userBrowser,
		"os":       userOS,
		"endpoint": filepath.Base(r.URL.Path),
	}
	host, err := os.Hostname()
	if err == nil {
		if idx := strings.IndexByte(host, '.'); idx > 0 {
			host = host[:idx]
		}
		stats["host"] = host
	}
	return stats
}

func parseUserAgent(uaString string) (browser, os string) {
	ua := uaparser.Parse(uaString)

	if ua.Browser != nil {
		browser = ua.Browser.Name
	}
	if ua.OS != nil {
		os = ua.OS.Name
	}

	return browser, os
}