package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *TheServiceCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list products\n"+
			"/get - get a entity\n"+
			"/delete — delete an existing entity\n"+
			"/new — create a new entity \n"+
			"/edit — edit a entity")

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Println("error send message %s", err)
		return
	}
}
