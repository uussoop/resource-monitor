package main

import (
	"os"
	"sync"

	"github.com/uussoop/resource-monitor/api"
)

func main() {
	os.Setenv("TELEGRAM_ADMINID", "")
	os.Setenv("TELEGRAM_APITOKEN", "")
	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		api.SimpleMonitoring(80, 80, 80)
	}()
	wg.Wait()

}
