package theService

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RatingTheServiceCommander) Delete(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		c.sendError("example: /delete__raiting__theservice 0", inputMsg.Chat.ID)
		return
	}

	rating, err := c.service.Describe(uint64(idx))
	if err != nil {
		c.sendError(fmt.Sprintf("fail to delete product: %v", err), inputMsg.Chat.ID)
		return
	}
	result, err := c.service.Remove(uint64(idx))
	if !result {
		c.sendError(fmt.Sprintf("fail to delete product: %v", err), inputMsg.Chat.ID)
	}
	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("delete product with idx %d\n %s", idx, rating),
	)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Println("error send message %s", err)
		return
	}
}
