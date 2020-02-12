package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	lez "github.com/NickBrisebois/BigLezChatBot/internal"
	"go/types"
	"log"
)

func main() {
	configPath := flag.String("config", "./config.toml", "Path to config.toml")
	flag.Parse()

	var config types.Config
	if _, err := toml.DecodeFile(*configPath, &config); err != nil {
		log.Fatal(err.Error())
	}

	bot.InitBot(&config)

}
