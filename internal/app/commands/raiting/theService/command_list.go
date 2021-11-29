package theService

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/real-mielofon/omp-bot/internal/app/path"
)

func (c *RatingTheServiceCommander) List(inputMsg *tgbotapi.Message) {
	outputMsgText := "Here all the ratings: \n\n"

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	ratings, err := c.rtgService.List(ctx, 0, itemsOnList)
	if err != nil {
		log.Printf("error c.service.List %s", err)
		return
	}

	for i, p := range ratings {
		outputMsgText += fmt.Sprintf("%3d: %s\n", i, p.String())
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)

	ctx, cancel = context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	if _, err := c.rtgService.Describe(ctx, itemsOnList); err == nil {
		serializedData, _ := json.Marshal(CallbackListData{
			Offset: itemsOnList,
		})

		callbackPath := path.CallbackPath{
			Domain:       "raiting",
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

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("error send message %s", err)
		return
	}
}
