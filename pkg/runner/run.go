package runner

import (
	"github.com/stokkelol/apod/pkg/nasa"
	"github.com/stokkelol/apod/pkg/telegram"
)

func Run(nasaKey, teleKey string, channelId int64) error {
	image, err := nasa.PullImage(nasaKey)
	if err != nil {
		return err
	}
	bot, err := telegram.NewApi(teleKey, channelId)
	if err != nil {
		return err
	}
	if err = bot.SendMessage(image); err != nil {
		return err
	}
	return nil
}
