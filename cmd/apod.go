package main

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/stokkelol/apod/pkg/runner"
)

var root = &cobra.Command{
	Use:   "apod",
	Short: "Publish latest astronomy picture of the day.",
	Long:  "Pull latest astronomy picture of the day (APOD) from NASA API and post it to Telegram channel",
	Run: func(cmd *cobra.Command, args []string) {
		nasaKey := os.Getenv("NASA_API_KEY")
		teleKey := os.Getenv("TELEGRAM_API_KEY")
		channelId, _ := strconv.Atoi(os.Getenv("TELEGRAM_CHANNEL_ID"))

		if err := runner.Run(nasaKey, teleKey, int64(channelId)); err != nil {
			log.Fatal(err.Error())
		}
	},
}

func main() {
	if err := root.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
