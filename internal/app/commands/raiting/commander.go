package raiting

import (
	"context"
	"github.com/real-mielofon/omp-bot/internal/pkg/logger"
	"github.com/real-mielofon/omp-bot/internal/service/raiting"

	"time"

	"github.com/real-mielofon/omp-bot/internal/app/commands/raiting/theService"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/real-mielofon/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(ctx context.Context, message *tgbotapi.Message, commandPath path.CommandPath)
}

type RatingCommander struct {
	bot              *tgbotapi.BotAPI
	serviceCommander Commander
}

func NewRaitingCommander(
	bot *tgbotapi.BotAPI,
	rtgService raiting.TheServiceService,
	timeout time.Duration,
) *RatingCommander {
	return &RatingCommander{
		bot:              bot,
		serviceCommander: theService.NewTheServiceCommander(bot, rtgService, timeout),
	}
}

func (c *RatingCommander) HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "service":
		c.serviceCommander.HandleCallback(ctx, callback, callbackPath)
	default:
		logger.InfoKV(ctx, "DemoCommander.HandleCallback: unknown subdomain", "Subdomain", callbackPath.Subdomain)
	}
}

func (c *RatingCommander) HandleCommand(ctx context.Context, msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "theservice":
		c.serviceCommander.HandleCommand(ctx, msg, commandPath)
	default:
		logger.InfoKV(ctx, "DemoCommander.HandleCommand: unknown subdomain", "Subdomain", commandPath.Subdomain)
	}
}
