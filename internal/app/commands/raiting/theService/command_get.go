package theService

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RatingTheServiceCommander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		c.sendError("example: /get__raiting__theservice 0", inputMsg.Chat.ID)
		return
	}

	rating, err := c.service.Describe(uint64(idx))
	if err != nil {
		c.sendError(fmt.Sprintf("fail to get rating with idx %d: %v", idx, err), inputMsg.Chat.ID)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		rating.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Println("error send message %s", err)
		return
	}
}
