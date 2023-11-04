package notification

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var BotToken string
var AdminId string
var AdminIdInt int64
var lastMessage time.Time
var Bot *tgbotapi.BotAPI

func SendTelegramMessage(message string) {

	if time.Since(lastMessage).Seconds() < 60 {
		return
	}
	Bot.Send(tgbotapi.NewMessage(AdminIdInt, message))
	lastMessage = time.Now()
}
