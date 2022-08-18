package tgservice

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"homework/config"
	"homework/internal/robot"
	"homework/internal/router"
	"homework/internal/storage/facade"
	"log"
)

var (
	helpCommand   = "help"
	listCommand   = "list"
	addCommand    = "add"
	removeCommand = "remove"
	updateCommand = "update"
)

type Service struct {
	*router.Router

	bot   *tgbotapi.BotAPI
	robot *robot.Robot
}

func NewService(storage facade.StorageFacade) (*Service, error) {
	bot, err := tgbotapi.NewBotAPI(config.ApiKey)
	if err != nil {
		return nil, errors.Wrap(err, "init tgbot")
	}
	bot.Debug = true

	r, err := robot.NewRobot(storage)
	if err != nil {
		return nil, errors.Wrap(err, "init robot")
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	service := &Service{bot: bot, robot: r}

	service.Router = router.New()
	service.AddCommand(helpCommand, service.Help)
	service.AddCommand(listCommand, service.List)
	service.AddCommand(addCommand, service.Add)
	service.AddCommand(updateCommand, service.Update)
	service.AddCommand(removeCommand, service.Remove)

	return service, nil
}

func (s *Service) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := s.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if cmd := update.Message.Command(); cmd != "" {
			res, err := s.Handle(cmd, update.Message.CommandArguments())
			if err != nil {
				msg.Text = errors.Wrap(err, "error while handle command").Error()
			} else {
				msg.Text = res
			}
		} else {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg.Text = fmt.Sprintf("you send <%v>", update.Message.Text)
		}
		_, err := s.bot.Send(msg)
		if err != nil {
			return errors.Wrap(err, "send tg message")
		}
	}
	return nil
}
