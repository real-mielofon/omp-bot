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

const itemsOnList = 10

func (c *TheServiceCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("error json.Unmarshal %s", parsedData)
	} else {

		outputMsgText := "Here all the raitings: \n\n"

		products := c.service.List()
		for i, p := range products[parsedData.Offset : parsedData.Offset+itemsOnList+1] {
			outputMsgText += fmt.Sprintf("%3d: %s\n", i+parsedData.Offset, p.String())
		}

		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)
		var buttons []tgbotapi.InlineKeyboardButton
		if parsedData.Offset-itemsOnList >= 0 {
			// Add Prev buttun
			serializedData, _ := json.Marshal(CallbackListData{
				Offset: parsedData.Offset - itemsOnList,
			})
			callbackPath := path.CallbackPath{
				Domain:       "rating",
				Subdomain:    "service",
				CallbackName: "list",
				CallbackData: string(serializedData),
			}

			buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Prev page", callbackPath.String()))
		}
		if parsedData.Offset+itemsOnList < len(products) {
			// Add Next buttun
			serializedData, _ := json.Marshal(CallbackListData{
				Offset: parsedData.Offset + itemsOnList,
			})
			callbackPath := path.CallbackPath{
				Domain:       "rating",
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

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Println("error send message %s", err)
			return
		}
		_, err = c.bot.AnswerCallbackQuery(tgbotapi.CallbackConfig{
			CallbackQueryID: callback.ID,
		})
		if err != nil {
			log.Println("error AnswerCallbackQuery %s", err)
			return
		}
	}
}
