package service

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/real-mielofon/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *RatingServiceCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("error json.Unmarshal %s", parsedData)
	} else {
		msg := tgbotapi.NewMessage(
			callback.Message.Chat.ID,
			fmt.Sprintf("Parsed: %+v\n", parsedData),
		)
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Println("error send message %s", err)
			return
		}
	}
}
