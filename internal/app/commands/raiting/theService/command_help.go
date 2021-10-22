package theService

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *RatingTheServiceCommander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"/help__raiting__theservice - help\n"+
			"/list__raiting__theservice - list products\n"+
			"/get__raiting__theservice - get a entity\n"+
			"/delete__raiting__theservice — delete an existing entity\n"+
			"/new__raiting__theservice — create a new entity \n"+
			"/edit__raiting__theservice — edit a entity")

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("error send message %s", err)
		return
	}
}
