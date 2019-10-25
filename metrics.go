package main

import "github.com/prometheus/client_golang/prometheus"

var baremetalHostsTotal = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "serverscom_baremetal_hosts_total",
})

var baremetalBalance = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "serverscom_baremetal_balance",
})

var baremetalEstimatedBalance = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "serverscom_baremetal_estimated_balance",
})

var baremetalHostUsageTraffic = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "serverscom_baremetal_host_usage_traffic",
}, []string{"host"})

var baremetalHostBillingPeriodTraffic = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "serverscom_baremetal_host_billing_period_traffic",
}, []string{"host"})

var baremetalHostPrice = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "serverscom_baremetal_host_price",
}, []string{"host"})

var baremetalHostDRACEnabled = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "serverscom_baremetal_host_drac_enabled",
}, []string{"host"})
