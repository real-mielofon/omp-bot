package rating

import (
	"log"

	"github.com/real-mielofon/omp-bot/internal/app/commands/rating/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/real-mielofon/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type RatingCommander struct {
	bot              *tgbotapi.BotAPI
	serviceCommander Commander
}

func NewRatingCommander(
	bot *tgbotapi.BotAPI,
) *RatingCommander {
	return &RatingCommander{
		bot:              bot,
		serviceCommander: service.NewRatingServiceCommander(bot),
	}
}

func (c *RatingCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "service":
		c.serviceCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("DemoCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *RatingCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "service":
		c.serviceCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("DemoCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
