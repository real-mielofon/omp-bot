package service

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/real-mielofon/omp-bot/internal/app/path"
	"github.com/real-mielofon/omp-bot/internal/service/raiting/service"
)

type RatingServiceCommander struct {
	bot            *tgbotapi.BotAPI
	serviceService *service.ServiceService
}

func NewRatingServiceCommander(
	bot *tgbotapi.BotAPI,
) *RatingServiceCommander {
	serviceService := service.NewService()

	return &RatingServiceCommander{
		bot:            bot,
		serviceService: serviceService,
	}
}

func (c *RatingServiceCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("RatingServiceCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *RatingServiceCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "edit":
		c.Edit(msg)
	case "new":
		c.New(msg)
	default:
		c.Default(msg)
	}
}

func (c *RatingServiceCommander) sendError(str string, inputMessageID int64) {
	log.Printf(str)
	msg := tgbotapi.NewMessage(
		inputMessageID,
		str,
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Println("error send message %s", err)
		return
	}
}
