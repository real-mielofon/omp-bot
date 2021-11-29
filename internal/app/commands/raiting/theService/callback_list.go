package theService

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/real-mielofon/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset uint64 `json:"offset"`
}

const itemsOnList = 10

func (c *RatingTheServiceCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("error json.Unmarshal %+v", parsedData)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	ratings, err := c.rtgService.List(ctx, parsedData.Offset, itemsOnList)
	if err != nil {
		c.sendError(fmt.Sprintf("error json.Unmarshal %+v", parsedData), callback.Message.Chat.ID)
		return
	}

	outputMsgText := "Here all the raitings: \n\n"
	for i, p := range ratings {
		outputMsgText += fmt.Sprintf("%3d: %s\n", uint64(i)+parsedData.Offset, p.String())
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)
	var buttons []tgbotapi.InlineKeyboardButton
	ctx, cancel = context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	if int(parsedData.Offset)-itemsOnList >= 0 {

		if _, err := c.rtgService.Describe(ctx, parsedData.Offset-itemsOnList); err == nil {
			// Add Prev buttun
			serializedData, _ := json.Marshal(CallbackListData{
				Offset: parsedData.Offset - itemsOnList,
			})
			callbackPath := path.CallbackPath{
				Domain:       "raiting",
				Subdomain:    "service",
				CallbackName: "list",
				CallbackData: string(serializedData),
			}

			buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Prev page", callbackPath.String()))
		}
	}
	if _, err := c.rtgService.Describe(ctx, parsedData.Offset+itemsOnList); err == nil {
		// Add Next buttun
		serializedData, _ := json.Marshal(CallbackListData{
			Offset: parsedData.Offset + itemsOnList,
		})
		callbackPath := path.CallbackPath{
			Domain:       "raiting",
			Subdomain:    "service",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()))
	}
	if len(buttons) > 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				buttons...,
			),
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("error send message %s", err)
		return
	}
	_, err = c.bot.AnswerCallbackQuery(tgbotapi.CallbackConfig{
		CallbackQueryID: callback.ID,
	})
	if err != nil {
		log.Printf("error AnswerCallbackQuery %s", err)
		return
	}
}
