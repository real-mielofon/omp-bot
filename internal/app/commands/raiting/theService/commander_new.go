package theService

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RatingTheServiceCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	helpNew := fmt.Sprintf("wrong args '%s'\n", args) +
		fmt.Sprintf("need: /new__raiting__theservice {ServiceID} {Value} {ReviewsCount}\n") +
		fmt.Sprintf("example: /new__raiting__theservice 8 5 0\n")
	if args == "" {
		c.sendError(helpNew, inputMessage.Chat.ID)
		return
	}

	parameters := strings.SplitN(args, " ", 3) //product.ServiceId product.Value product.ReviewsCount
	if len(parameters) != 3 {
		log.Println("wrong args", args)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			helpNew,
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Println("error send message %s", err)
			return
		}
		return
	}
	productServiceId, err := strconv.Atoi(parameters[0])
	if err != nil {
		c.sendError(fmt.Sprintf("error ServiceId value: %s", parameters[0]), inputMessage.Chat.ID)
		return
	}
	productValue, err := strconv.Atoi(parameters[1])
	if err != nil {
		c.sendError(fmt.Sprintf("error Value value: %s", parameters[1]), inputMessage.Chat.ID)
		return
	}
	productReviewsCount, err := strconv.Atoi(parameters[2])
	if err != nil {
		c.sendError(fmt.Sprintf("error ReviewsCount value: %s", parameters[2]), inputMessage.Chat.ID)
		return
	}

	product, err := c.service.New()
	if err != nil {
		log.Printf("error new product %v", err)
		return
	}
	product.ServiceId = productServiceId
	product.Value = productValue
	product.ReviewsCount = productReviewsCount
	product.UpdatedTs = time.Unix(int64(inputMessage.Date), 0)

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
