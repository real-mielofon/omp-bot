package theService

import (
	"context"
	"fmt"
	"github.com/real-mielofon/omp-bot/internal/model/raiting"
	"github.com/real-mielofon/omp-bot/internal/pkg/logger"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RatingTheServiceCommander) Edit(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	helpEdit := fmt.Sprintf("wrong args '%s'\n", args) +
		"need: /edit__raiting__thrservice {idx} {ServiceID} {Value} {ReviewsCount}\n" +
		"example: /edit__raiting__theservice 5 8 5 0\n"

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
	service := raiting.TheService{}
	id, err := strconv.Atoi(parameters[1])
	if err != nil {
		c.sendError(fmt.Sprintf("error ServiceId value: %s", parameters[1]), inputMsg.Chat.ID)
		return
	}
	if id < 0 {
		c.sendError(fmt.Sprintf("error ServiceId value: %s", parameters[1]), inputMsg.Chat.ID)
		return
	}
	service.ID = uint64(id)

	service.Value, err = strconv.Atoi(parameters[2])
	if err != nil {
		c.sendError(fmt.Sprintf("error Value value: %s", parameters[2]), inputMsg.Chat.ID)
		return
	}
	service.ReviewsCount, err = strconv.Atoi(parameters[3])
	if err != nil {
		c.sendError(fmt.Sprintf("error ReviewsCount value: %s", parameters[3]), inputMsg.Chat.ID)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	err = c.rtgService.Update(ctx, uint64(idx), service)
	if err != nil {
		c.sendError(fmt.Sprintf("fail to get service with idx %d: %v", idx, err), inputMsg.Chat.ID)
		return
	}
	service.UpdatedTs = time.Unix(int64(inputMsg.Date), 0)

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("edit service idx: %d\n%s", idx, service.String()),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		logger.ErrorKV(ctx, "error send message", "err", err)
		return
	}
}
