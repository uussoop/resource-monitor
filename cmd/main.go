package main

import (
	"os"

	"github.com/uussoop/resource-monitor/api"
)

func main() {
	os.Setenv("TELEGRAM_ADMINID", "")
	os.Setenv("TELEGRAM_APITOKEN", "")
	go api.SimpleMonitoring(80, 80, 80)

	for {
	}
}
