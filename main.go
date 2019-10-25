package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/bidease/scomportal"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var confFile = flag.String("conf", "/etc/scom_exporter.yml", "path to config file")

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	flag.Parse()
	conf.Read(*confFile)

	prometheus.MustRegister(
		baremetalHostsTotal,
		baremetalBalance,
		baremetalEstimatedBalance,
		baremetalHostUsageTraffic,
		baremetalHostBillingPeriodTraffic,
		baremetalHostPrice,
		baremetalHostDRACEnabled,
	)
}

func main() {
	go getMetrics()

	http.Handle(conf.MetricsEndpoint, promhttp.Handler())
	log.Fatal(http.ListenAndServe(conf.WebMetricsAddr, nil))
}

func getMetrics() {
	api := scomportal.NewAPI(conf.Auth.Email, conf.Auth.Token)

	for {
		go getBaremetalHostsTotal(api)
		go getBaremetalBalance(api)
		go getHostMetrics(api)

		time.Sleep(time.Duration(time.Second * time.Duration(conf.ScrapeInterval)))
	}
}
