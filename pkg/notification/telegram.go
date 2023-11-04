package notification

import (
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var BotToken string
var AdminId string
var lastMessage time.Time
var Bot *tgbotapi.BotAPI

func SendTelegramMessage(message string) {

	adminIdInt, err := (strconv.Atoi(AdminId))

	if err != nil {
		panic(err)
	}
	if time.Since(lastMessage).Seconds() < 60 {
		return
	}
	Bot.Send(tgbotapi.NewMessage(int64(adminIdInt), message))
	lastMessage = time.Now()
}
