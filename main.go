package main

import (
	"github.com/stokkelol/apod/pkg/runner"
	"log"
	"os"
	"strconv"
)

func main() {
	nasaKey := os.Getenv("NASA_API_KEY")
	teleKey := os.Getenv("TELEGRAM_API_KEY")
	channelId, _ := strconv.Atoi(os.Getenv("TELEGRAM_CHANNEL_ID"))

	if err := runner.Run(nasaKey, teleKey, int64(channelId)); err != nil {
		log.Fatal(err.Error())
	}
}
