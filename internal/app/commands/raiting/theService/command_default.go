package theService

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/real-mielofon/omp-bot/internal/pkg/logger"
)

func (c *RatingTheServiceCommander) Default(ctx context.Context, inputMsg *tgbotapi.Message) {
	logger.InfoKV(ctx, "default", "UserName", inputMsg.From.UserName, "Text", inputMsg.Text)

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "You wrote: "+inputMsg.Text)

	_, err := c.bot.Send(msg)
	if err != nil {
		logger.ErrorKV(ctx, "error send message", "err", err)
		return
	}
}
