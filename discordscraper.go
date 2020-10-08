package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	scraper "github.com/NickBrisebois/DiscordServerMessageScraper/scraper"
	"log"
    "fmt"
)

var version string
var gitCommit string
var buildTime string

func main() {
	configPath := flag.String("config", "./config.toml", "Path to config.toml")
    versionFlag := flag.Bool("v", false, "Show version")
	flag.Parse()

    if *versionFlag {
        fmt.Println("Version: " + version)
        fmt.Println("Commit: " + gitCommit)
        fmt.Println("Built: " + buildTime)
        return;
    }

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
