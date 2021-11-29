package theService

import (
	"github.com/real-mielofon/omp-bot/internal/service/raiting"
	"log"
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

func (c *RatingTheServiceCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("RatingTheServiceCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *RatingTheServiceCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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

func (c *RatingTheServiceCommander) sendError(str string, inputMessageID int64) {
	log.Println(str)
	msg := tgbotapi.NewMessage(
		inputMessageID,
		str,
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("error send message %s", err)
		return
	}
}
