package main

//This script listens on a given TCP port for
//HTTP REST Get messages than scraps the given
//Open vSwtich entry and gives back the stats
//in Prometheus compatible format

//Written by Megyo @ LeanNet

import (
	"log"
	"net/http"
	"ovs/ovs"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//the TCP port that this scripts listens
var listenPort string = ":8081"

func handler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("target")
	if target == "" {
		http.Error(w, "Bad request!\nCorrect format is: http://"+r.Host+"?target=<targetIP>\nfor example:http://127.0.0.1:8081/metrics?target=127.0.0.1", 400)
		return
	}
	c := OvsPromCollector{
		ip:        target,
		port:      ovs.OvsDefaultPort,
		ovsReader: ovs.CliDumpReader,
	}
	registry := prometheus.NewRegistry()
	registry.MustRegister(c)
	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/metrics", handler)
	log.Fatal(http.ListenAndServe(listenPort, nil))
}
