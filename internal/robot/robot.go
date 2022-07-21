package robot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"homework/config"
	"homework/internal/router"
	"homework/internal/storage"
	"log"
)

type Robot struct {
	bot    *tgbotapi.BotAPI
	storage *storage.Storage
}

func Init() (*Robot, error) {
	bot, err := tgbotapi.NewBotAPI(config.ApiKey)
	if err != nil {
		return nil, errors.Wrap(err, "init tgbot")
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Robot{bot: bot, storage: storage.NewStorage()}, nil
}

func (r *Robot) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := r.bot.GetUpdatesChan(u)

	router := router.New()
	router.Add(helpCommand, r.helpFunc)
	router.Add(listCommand, r.listFunc)
	router.Add(addCommand, r.addFunc)
	router.Add(removeCommand, r.removeFunc)
	router.Add(updateCommand, r.updateFunc)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if cmd := update.Message.Command(); cmd != "" {
			msg.Text = router.Handle(cmd, update.Message.CommandArguments())
		} else {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg.Text = fmt.Sprintf("you send <%v>", update.Message.Text)
		}

		_, err := r.bot.Send(msg)
		if err != nil {
			return errors.Wrap(err, "send tg message")
		}
	}
	return nil
}
