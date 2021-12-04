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

func (c *RatingTheServiceCommander) New(ctx context.Context, inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	helpNew := fmt.Sprintf("wrong args '%s'\n", args) +
		"need: /new__raiting__theservice {ServiceID} {Value} {ReviewsCount}\n" +
		"example: /new__raiting__theservice 8 5 0\n"
	if args == "" {
		c.sendError(ctx, helpNew, inputMsg.Chat.ID)
		return
	}

	parameters := strings.SplitN(args, " ", 3) //product.ServiceId product.Value product.ReviewsCount
	if len(parameters) != 3 {
		logger.InfoKV(ctx, "wrong args", "args", args)

		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			helpNew,
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			logger.ErrorKV(ctx, "error send message", "err", err)
			return
		}
		return
	}

	service := raiting.TheService{}
	var err error

	id, err := strconv.Atoi(parameters[0])
	if err != nil {
		c.sendError(ctx, fmt.Sprintf("error ServiceId value: %s", parameters[1]), inputMsg.Chat.ID)
		return
	}
	if id < 0 {
		c.sendError(ctx, fmt.Sprintf("error ServiceId value: %s", parameters[1]), inputMsg.Chat.ID)
		return
	}
	service.ID = uint64(id)

	service.Value, err = strconv.Atoi(parameters[1])

	if err != nil {
		c.sendError(ctx, fmt.Sprintf("error Value value: %s", parameters[1]), inputMsg.Chat.ID)
		return
	}
	service.ReviewsCount, err = strconv.Atoi(parameters[2])
	if err != nil {
		c.sendError(ctx, fmt.Sprintf("error ReviewsCount value: %s", parameters[2]), inputMsg.Chat.ID)
		return
	}

	service.UpdatedTs = time.Unix(int64(inputMsg.Date), 0)

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	idx, err := c.rtgService.Create(ctx, service)
	if err != nil {
		c.sendError(ctx, fmt.Sprintf("error new product %v", err), inputMsg.Chat.ID)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("%d: %s", idx, service.String()),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		logger.ErrorKV(ctx, "error send message", "err", err)
		return
	}
}
