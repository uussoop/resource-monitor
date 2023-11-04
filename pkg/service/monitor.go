package service

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/uussoop/resource-monitor/pkg/notification"
	resourceapi "github.com/uussoop/resource-monitor/pkg/resourceApi"
)

func Monitor(cpuTop int, memTop int, diskTop int) {
	notification.Bot.Send(tgbotapi.NewMessage(notification.AdminIdInt, "Starting monitoring"))
	for {

		cpuUtilization, _ := resourceapi.GetCpuUtilization()
		if cpuUtilization > cpuTop {
			notification.SendTelegramMessage("CPU is reaching " + fmt.Sprint(cpuUtilization) + "%")
		}
		memoryUsage := resourceapi.GetMemoryUsage()
		if memoryUsage > memTop {
			notification.SendTelegramMessage("Memory is reaching " + fmt.Sprint(memoryUsage) + "%")
		}
		diskUsage := resourceapi.CheckIfAnyDeviceReachingQouta(diskTop)
		if diskUsage {
			notification.SendTelegramMessage("Disk is reaching " + fmt.Sprint(diskTop) + "%")
		}
		//sleep
		time.Sleep(100 * time.Millisecond)
	}
}
