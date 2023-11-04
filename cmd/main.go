package main

import (
	"fmt"
	"sync"

	"github.com/uussoop/resource-monitor/api"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Starting monitoring")

		api.SimpleMonitoring(80, 95, 80)
	}(&wg)

	wg.Wait()
}
