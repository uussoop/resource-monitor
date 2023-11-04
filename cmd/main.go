package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/uussoop/resource-monitor/api"
)

func main() {
	os.Setenv("TELEGRAM_ADMINID", "")
	os.Setenv("TELEGRAM_APITOKEN", "")

	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Starting monitoring")

		api.SimpleMonitoring(80, 95, 80)
	}(&wg)

	wg.Wait()
}
