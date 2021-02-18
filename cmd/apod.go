package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stokkelol/apod/pkg/nasa"
	"github.com/stokkelol/apod/pkg/telegram"
	"log"
	"os"
	"strconv"
)

var root = &cobra.Command{
	Use:   "apod",
	Short: "Publish latest astronomy picture of the day.",
	Long:  "Pull latest astronomy picture of the day (APOD) from NASA API and post it to Telegram channel",
	Run: func(cmd *cobra.Command, args []string) {
		nasaKey := os.Getenv("NASA_API_KEY")
		teleKey := os.Getenv("TELEGRAM_API_KEY")
		channelId, _ := strconv.Atoi(os.Getenv("TELEGRAM_CHANNEL_ID"))

		image, err := nasa.PullImage(nasaKey)
		if err != nil {
			log.Fatal(err.Error())
		}
		bot, err := telegram.NewApi(teleKey, int64(channelId))
		if err != nil {
			log.Fatal(err.Error())
		}
		if err := bot.SendMessage(image); err != nil {
			log.Fatal(err.Error())
		}
	},
}

func Execute() error {
	return root.Execute()
}
