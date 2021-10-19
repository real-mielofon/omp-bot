package theService

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RatingTheServiceCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	helpEdit := fmt.Sprintf("wrong args '%s'\n", args) +
		fmt.Sprintf("need: /edit__raiting__thrservice {idx} {ServiceID} {Value} {ReviewsCount}\n") +
		fmt.Sprintf("example: /edit__raiting__theservice 5 8 5 0\n")

	if args == "" {
		c.sendError(helpEdit, inputMessage.Chat.ID)
		return
	}
	parameters := strings.SplitN(args, " ", 4) //product.ServiceId product.Value product.ReviewsCount
	if len(parameters) != 4 {
		c.sendError(helpEdit, inputMessage.Chat.ID)
		return
	}

	idx, err := strconv.Atoi(parameters[0])
	if err != nil {
		c.sendError(fmt.Sprintf("error idx value: %s", parameters[0]), inputMessage.Chat.ID)
		return
	}
	productServiceId, err := strconv.Atoi(parameters[1])
	if err != nil {
		c.sendError(fmt.Sprintf("error ServiceId value: %s", parameters[1]), inputMessage.Chat.ID)
		return
	}
	productValue, err := strconv.Atoi(parameters[2])
	if err != nil {
		c.sendError(fmt.Sprintf("error Value value: %s", parameters[2]), inputMessage.Chat.ID)
		return
	}
	productReviewsCount, err := strconv.Atoi(parameters[3])
	if err != nil {
		c.sendError(fmt.Sprintf("error ReviewsCount value: %s", parameters[3]), inputMessage.Chat.ID)
		return
	}

	product, err := c.service.Get(idx)
	if err != nil {
		log.Printf("fail to get service with idx %d: %v", idx, err)
		return
	}
	product.ServiceId = productServiceId
	product.Value = productValue
	product.ReviewsCount = productReviewsCount
	product.UpdatedTs = time.Unix(int64(inputMessage.Date), 0)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("edit service idx: %d\n%s", idx, product.String()),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Println("error send message %s", err)
		return
	}
}
