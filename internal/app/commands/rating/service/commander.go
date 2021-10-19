package service

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/real-mielofon/omp-bot/internal/app/path"
	"github.com/real-mielofon/omp-bot/internal/service/raiting/service"
)

type ServiceCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type TheServiceCommander struct {
	bot     *tgbotapi.BotAPI
	service *service.Service
}

func NewTheServiceCommander(
	bot *tgbotapi.BotAPI,
) *TheServiceCommander {
	serviceService := service.NewService()

	return &TheServiceCommander{
		bot:     bot,
		service: serviceService,
	}
}

func (c *TheServiceCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("TheServiceCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *TheServiceCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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

func (c *TheServiceCommander) sendError(str string, inputMessageID int64) {
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
