package api

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/uussoop/resource-monitor/pkg/notification"
	"github.com/uussoop/resource-monitor/pkg/service"
)

func SimpleMonitoring(cpu int, memory int, disk int) {
	notification.AdminId = os.Getenv("TELEGRAM_ADMINID")
	notification.BotToken = os.Getenv("TELEGRAM_APITOKEN")
	if notification.AdminId == "" {
		panic("TELEGRAM_ADMINID is not set")
	}
	if notification.BotToken == "" {
		panic("TELEGRAM_APITOKEN is not set")
	}
	var err error
	notification.Bot, err = tgbotapi.NewBotAPI(notification.BotToken)
	if err != nil {
		panic(err)
	}
	service.Monitor(cpu, memory, disk)
}
