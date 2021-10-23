package theService

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/service/raiting/theService"
	"log"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RatingTheServiceCommander) New(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	helpNew := fmt.Sprintf("wrong args '%s'\n", args) +
		"need: /new__raiting__theservice {ServiceID} {Value} {ReviewsCount}\n" +
		"example: /new__raiting__theservice 8 5 0\n"
	if args == "" {
		c.sendError(helpNew, inputMsg.Chat.ID)
		return
	}

	parameters := strings.SplitN(args, " ", 3) //product.ServiceId product.Value product.ReviewsCount
	if len(parameters) != 3 {
		log.Println("wrong args", args)
		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			helpNew,
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("error send message %s", err)
			return
		}
		return
	}

	raiting := theService.TheService{}
	var err error
	raiting.ServiceId, err = strconv.Atoi(parameters[0])
	if err != nil {
		c.sendError(fmt.Sprintf("error ServiceId value: %s", parameters[0]), inputMsg.Chat.ID)
		return
	}
	raiting.Value, err = strconv.Atoi(parameters[1])
	if err != nil {
		c.sendError(fmt.Sprintf("error Value value: %s", parameters[1]), inputMsg.Chat.ID)
		return
	}
	raiting.ReviewsCount, err = strconv.Atoi(parameters[2])
	if err != nil {
		c.sendError(fmt.Sprintf("error ReviewsCount value: %s", parameters[2]), inputMsg.Chat.ID)
		return
	}

	raiting.UpdatedTs = time.Unix(int64(inputMsg.Date), 0)

	idx, err := c.service.Create(raiting)
	if err != nil {
		c.sendError(fmt.Sprintf("error new product %v", err), inputMsg.Chat.ID)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("%d: %s", idx, raiting.String()),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("error send message %s", err)
		return
	}
}
