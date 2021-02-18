package telegram

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/stokkelol/apod/pkg/nasa"
)

type Bot struct {
	bot    *telegram.BotAPI
	token  string
	chatID int64
}

func NewApi(token string, chatID int64) (*Bot, error) {
	bot, err := telegram.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		bot:    bot,
		token:  token,
		chatID: chatID,
	}, nil
}

func (b *Bot) SendMessage(image *nasa.ApodImage) error {
	bc := telegram.BaseChat{
		ChatID: b.chatID,
	}
	_, err := b.bot.Send(telegram.MessageConfig{
		BaseChat: bc,
		Text:     image.Title,
	})
	if err != nil {
		return err
	}
	_, err = b.bot.Send(telegram.MessageConfig{
		BaseChat: bc,
		Text:     image.Url,
	})
	if err != nil {
		return err
	}
	_, err = b.bot.Send(telegram.MessageConfig{
		BaseChat: bc,
		Text:     image.Explanation,
	})
	if err != nil {
		return err
	}

	return nil
}
