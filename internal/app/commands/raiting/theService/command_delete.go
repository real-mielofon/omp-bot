package theService

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RatingTheServiceCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		c.sendError("example: /delete__raiting__theservice 0", inputMessage.Chat.ID)
		return
	}

	product, err := c.service.Get(idx)
	if err != nil {
		c.sendError(fmt.Sprintf("fail to delete product: %v", err), inputMessage.Chat.ID)
		return
	}
	err = c.service.Delete(idx)

	if err != nil {
		c.sendError(fmt.Sprintf("fail to delete product: %v", err), inputMessage.Chat.ID)
	}
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("delete product with idx %d\n %s", idx, product),
	)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Println("error send message %s", err)
		return
	}
}
