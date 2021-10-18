package service

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RatingServiceCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	err = c.serviceService.Delete(idx)

	msg := tgbotapi.MessageConfig{}
	if err != nil {
		log.Printf("fail to delete product with idx %d: %v", idx, err)
		msg = tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("fail to delete product with idx %d: %v", idx, err),
		)
	} else {
		msg = tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("delete product with idx %d: %v", idx, err),
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Println("error send message %s", err)
		return
	}
}
