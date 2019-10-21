package main

import (
	"log"
	"strconv"

	"github.com/bidease/scomportal"
)

func getBaremetalHostsTotal(api *scomportal.API) {
	if equipment, err := api.GetBaremetalEquipment(); err == nil {
		baremetalHostsTotal.Set(float64(equipment.Data.Equipment.Hosts))
	} else {
		log.Println(err)
	}
}

func getBaremetalBalance(api *scomportal.API) {
	balances, err := api.GetBaremetalBalance()
	if err != nil {
		log.Fatalln(err)
	}

	balance, _ := strconv.ParseFloat(balances.Data.Balance, 64)
	baremetalBalance.Set(balance)
	estimatedBalance, _ := strconv.ParseFloat(balances.Data.EstimatedBalance, 64)
	baremetalEstimatedBalance.Set(estimatedBalance)
}

func getTrafficHost(api *scomportal.API) {
	hosts, err := api.GetBaremetalHosts()
	if err != nil {
		log.Fatalln(err)
	}

	for _, host := range hosts.Data {
		traffic, err := api.GetBaremetalHostTraffic(host.ID)
		if err != nil {
			log.Println(err)
		}

		usageTraffic, _ := strconv.ParseFloat(traffic.Data.Commit.UsageQuantity, 64)
		baremetalHostUsageTraffic.WithLabelValues(host.Title).Set(usageTraffic)
		baremetalHostBillingPeriodTraffic.WithLabelValues(host.Title).Set(traffic.Data.Commit.CommitValueForBillingPeriod)
	}
}

func getBaremetalHostPrice(api *scomportal.API) {
	hosts, err := api.GetBaremetalHosts()
	if err != nil {
		log.Fatalln(err)
	}

	for _, host := range hosts.Data {
		services, err := api.GetServices(host.ID)
		if err != nil {
			log.Println(err)
		}

		var price float64

		for _, v := range services.Data {
			price = price + v.Price
		}

		baremetalHostPrice.WithLabelValues(host.Title).Set(price)
	}
}