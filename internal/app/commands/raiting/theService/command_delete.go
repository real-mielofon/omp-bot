package theService

import (
	"context"
	"fmt"
	"github.com/real-mielofon/omp-bot/internal/pkg/logger"

	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RatingTheServiceCommander) Delete(ctx context.Context, inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		c.sendError(ctx, "example: /delete__raiting__theservice 0", inputMsg.Chat.ID)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	rating, err := c.rtgService.Describe(ctx, uint64(idx))
	if err != nil {
		c.sendError(ctx, fmt.Sprintf("fail to delete product: %v", err), inputMsg.Chat.ID)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	result, err := c.rtgService.Remove(ctx, uint64(idx))
	if !result {
		c.sendError(ctx, fmt.Sprintf("fail to delete product: %v", err), inputMsg.Chat.ID)
	}
	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("delete product with idx %d\n %s", idx, rating),
	)
	_, err = c.bot.Send(msg)
	if err != nil {
		logger.ErrorKV(ctx, "error send message", "err", err)
		return
	}
}
