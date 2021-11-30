package router

import (
	"context"
	"github.com/real-mielofon/omp-bot/internal/pkg/logger"
	"github.com/real-mielofon/omp-bot/internal/service/raiting"

	"runtime/debug"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	commandRaiting "github.com/real-mielofon/omp-bot/internal/app/commands/raiting"
	"github.com/real-mielofon/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(ctx context.Context, callback *tgbotapi.Message, commandPath path.CommandPath)
}

type Router struct {
	// bot
	bot *tgbotapi.BotAPI

	// demoCommander
	demoCommander Commander
	// user
	// access
	// buy
	// delivery
	// recommendation
	// travel
	// loyalty
	// bank
	// subscription
	// license
	// insurance
	// payment
	// storage
	// streaming
	// business
	// work
	// service
	// exchange
	// estate
	// raiting
	raitingCommander Commander
	// security
	// cinema
	// logistic
	// product
	// education
}

func NewRouter(
	bot *tgbotapi.BotAPI,
	rtgService raiting.TheServiceService,
	timeOut time.Duration,
) *Router {
	return &Router{
		// bot
		bot: bot,
		// demoCommander
		// user
		// access
		// buy
		// delivery
		// recommendation
		// travel
		// loyalty
		// bank
		// subscription
		// license
		// insurance
		// payment
		// storage
		// streaming
		// business
		// work
		// service
		// exchange
		// estate
		// raiting
		raitingCommander: commandRaiting.NewRaitingCommander(bot, rtgService, timeOut),
		// security
		// cinema
		// logistic
		// product
		// education
	}
}

func (c *Router) HandleUpdate(ctx context.Context, update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			logger.InfoKV(ctx, "recovered from panic",
				"panicValue", panicValue,
				"stack", string(debug.Stack()))
		}
	}()

	switch {
	case update.CallbackQuery != nil:
		c.handleCallback(ctx, update.CallbackQuery)
	case update.Message != nil:
		c.handleMessage(ctx, update.Message)
	}
}

func (c *Router) handleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery) {
	callbackPath, err := path.ParseCallback(callback.Data)
	if err != nil {
		logger.ErrorKV(ctx, "Router.handleCallback: error parsing callback data", "callback.Data", callback.Data, "err", err)
		return
	}

	switch callbackPath.Domain {
	case "demo":
		c.demoCommander.HandleCallback(ctx, callback, callbackPath)
	case "user":
		break
	case "access":
		break
	case "buy":
		break
	case "delivery":
		break
	case "recommendation":
		break
	case "travel":
		break
	case "loyalty":
		break
	case "bank":
		break
	case "subscription":
		break
	case "license":
		break
	case "insurance":
		break
	case "payment":
		break
	case "storage":
		break
	case "streaming":
		break
	case "business":
		break
	case "work":
		break
	case "service":
		break
	case "exchange":
		break
	case "estate":
		break
	case "raiting":
		c.raitingCommander.HandleCallback(ctx, callback, callbackPath)
	case "security":
		break
	case "cinema":
		break
	case "logistic":
		break
	case "product":
		break
	case "education":
		break
	default:
		logger.InfoKV(ctx, "Router.handleCallback: unknown domain", "domain", callbackPath.Domain)
	}
}

func (c *Router) handleMessage(ctx context.Context, msg *tgbotapi.Message) {
	if !msg.IsCommand() {
		c.showCommandFormat(ctx, msg)

		return
	}

	commandPath, err := path.ParseCommand(msg.Command())
	if err != nil {
		logger.ErrorKV(ctx, "Router.handleCallback: error parsing callback data",
			"command", msg.Command(),
			"err", err)
		return
	}

	switch commandPath.Domain {
	case "demo":
		c.demoCommander.HandleCommand(ctx, msg, commandPath)
	case "user":
		break
	case "access":
		break
	case "buy":
		break
	case "delivery":
		break
	case "recommendation":
		break
	case "travel":
		break
	case "loyalty":
		break
	case "bank":
		break
	case "subscription":
		break
	case "license":
		break
	case "insurance":
		break
	case "payment":
		break
	case "storage":
		break
	case "streaming":
		break
	case "business":
		break
	case "work":
		break
	case "service":
		break
	case "exchange":
		break
	case "estate":
		break
	case "raiting":
		c.raitingCommander.HandleCommand(ctx, msg, commandPath)
	case "security":
		break
	case "cinema":
		break
	case "logistic":
		break
	case "product":
		break
	case "education":
		break
	default:
		logger.InfoKV(ctx, "Router.handleCallback: unknown domain", "domain", commandPath.Domain)
	}
}

func (c *Router) showCommandFormat(ctx context.Context, inputMessage *tgbotapi.Message) {
	outputMsg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Command format: /{command}__{domain}__{subdomain}")

	_, err := c.bot.Send(outputMsg)
	if err != nil {
		logger.ErrorKV(ctx, "Router.showCommandFormat: error sending reply message to chat", "err", err)
	}
}
