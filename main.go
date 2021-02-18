package main

import (
	"github.com/stokkelol/apod/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("error while running APOD command: %s", err.Error())
	}
}
