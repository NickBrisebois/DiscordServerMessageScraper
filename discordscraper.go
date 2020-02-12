package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	scraper "github.com/NickBrisebois/DiscordServerMessageScraper/scraper"
	"log"
)

func main() {
	configPath := flag.String("config", "./config.toml", "Path to config.toml")
	flag.Parse()

	var config scraper.Config
	if _, err := toml.DecodeFile(*configPath, &config); err != nil {
		log.Fatal(err.Error())
	}

	scraper := scraper.NewServerScraper(&config)
	err := scraper.InitScraper()

	if err != nil {
		log.Fatalf(err.Error())
	}

}
