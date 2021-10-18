package service

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/real-mielofon/omp-bot/internal/app/path"
)

func (c *RatingServiceCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the raitings: \n\n"

	products := c.serviceService.List()
	for i, p := range products[:itemsOnList+1] {
		outputMsgText += fmt.Sprintf("%3d: %s\n", i, p.String())
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	if len(products) > itemsOnList {
		serializedData, _ := json.Marshal(CallbackListData{
			Offset: itemsOnList + 1,
		})

		callbackPath := path.CallbackPath{
			Domain:       "rating",
			Subdomain:    "service",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
			),
		)
	}

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Println("error send message %s", err)
		return
	}
}
