package notification

import (
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var BotToken string
var AdminId string
var lastMessage time.Time

func SendTelegramMessage(message string) {

	adminIdInt, err := (strconv.Atoi(AdminId))

	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		panic(err)
	}
	if time.Since(lastMessage).Seconds() < 60 {
		return
	}
	bot.Send(tgbotapi.NewMessage(int64(adminIdInt), message))
	lastMessage = time.Now()
}