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

func (c *RatingTheServiceCommander) Edit(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	helpEdit := fmt.Sprintf("wrong args '%s'\n", args) +
		fmt.Sprintf("need: /edit__raiting__thrservice {idx} {ServiceID} {Value} {ReviewsCount}\n") +
		fmt.Sprintf("example: /edit__raiting__theservice 5 8 5 0\n")

	if args == "" {
		c.sendError(helpEdit, inputMsg.Chat.ID)
		return
	}
	parameters := strings.SplitN(args, " ", 4) //product.ServiceId product.Value product.ReviewsCount
	if len(parameters) != 4 {
		c.sendError(helpEdit, inputMsg.Chat.ID)
		return
	}

	idx, err := strconv.Atoi(parameters[0])
	if err != nil {
		c.sendError(fmt.Sprintf("error idx value: %s", parameters[0]), inputMsg.Chat.ID)
		return
	}
	raiting := theService.TheService{}
	raiting.ServiceId, err = strconv.Atoi(parameters[1])
	if err != nil {
		c.sendError(fmt.Sprintf("error ServiceId value: %s", parameters[1]), inputMsg.Chat.ID)
		return
	}
	raiting.Value, err = strconv.Atoi(parameters[2])
	if err != nil {
		c.sendError(fmt.Sprintf("error Value value: %s", parameters[2]), inputMsg.Chat.ID)
		return
	}
	raiting.ReviewsCount, err = strconv.Atoi(parameters[3])
	if err != nil {
		c.sendError(fmt.Sprintf("error ReviewsCount value: %s", parameters[3]), inputMsg.Chat.ID)
		return
	}

	err = c.service.Update(uint64(idx), raiting)
	if err != nil {
		c.sendError(fmt.Sprintf("fail to get service with idx %d: %v", idx, err), inputMsg.Chat.ID)
		return
	}
	raiting.UpdatedTs = time.Unix(int64(inputMsg.Date), 0)

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("edit service idx: %d\n%s", idx, raiting.String()),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Println("error send message %s", err)
		return
	}
}
