package theService

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RatingTheServiceCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		c.sendError("example: /get__raiting__theservice 0", inputMessage.Chat.ID)
		return
	}

	product, err := c.service.Get(idx)
	if err != nil {
		log.Printf("fail to get product with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		product.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Println("error send message %s", err)
		return
	}
}
