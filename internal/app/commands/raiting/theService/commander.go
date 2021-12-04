package theService

import (
	"context"
	"github.com/real-mielofon/omp-bot/internal/pkg/logger"
	"github.com/real-mielofon/omp-bot/internal/service/raiting"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/real-mielofon/omp-bot/internal/app/path"
)

type RatingTheServiceCommander struct {
	bot        *tgbotapi.BotAPI
	rtgService raiting.TheServiceService
	timeout    time.Duration
}

func NewTheServiceCommander(
	bot *tgbotapi.BotAPI,
	rtgService raiting.TheServiceService,
	timeout time.Duration,
) *RatingTheServiceCommander {

	return &RatingTheServiceCommander{
		bot:        bot,
		rtgService: rtgService,
		timeout:    timeout,
	}
}

func (c *RatingTheServiceCommander) HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(ctx, callback, callbackPath)
	default:
		logger.InfoKV(ctx, "RatingTheServiceCommander.HandleCallback: unknown callback name", "CallbackName", callbackPath.CallbackName)
	}
}

func (c *RatingTheServiceCommander) HandleCommand(ctx context.Context, msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(ctx, msg)
	case "list":
		c.List(ctx, msg)
	case "get":
		c.Get(ctx, msg)
	case "delete":
		c.Delete(ctx, msg)
	case "edit":
		c.Edit(ctx, msg)
	case "new":
		c.New(ctx, msg)
	default:
		c.Default(ctx, msg)
	}
}

func (c *RatingTheServiceCommander) sendError(ctx context.Context, str string, inputMessageID int64) {
	logger.InfoKV(ctx, "Error", "str", str)
	msg := tgbotapi.NewMessage(
		inputMessageID,
		str,
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		logger.ErrorKV(ctx, "error send message", "err", err)
		return
	}
}
