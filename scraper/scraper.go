package biglez_scraper

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

type ServerScraper struct {
	botConf *Config
	sesh *discordgo.Session
}

func NewServerScraper(config *Config) *ServerScraper {
	return &ServerScraper{
		botConf: config,
	}
}

func (sc *ServerScraper) InitScraper() error {
	log.Println("Initializing The Big LezTM Server Scraper")

	var err error
	sc.sesh, err = discordgo.New("Bot " + sc.botConf.DiscordToken)

	if err != nil {
		return err
	}

	err = sc.sesh.Open()
	if err != nil {
		return err
	}

	for _, guild := range sc.sesh.State.Guilds {
		channels, _ := sc.sesh.GuildChannels(guild.ID)
		fmt.Println(channels)
	}

	//sc.BulkDownloadMessages()

	sc.sesh.Close()
	return nil
}

func (sc *ServerScraper) BulkDownloadMessages() {
	startID := sc.botConf.StartMessageID
	var messages []*discordgo.Message
	var err error

	for {
		blankLineCount := 0
		messages, err = sc.sesh.ChannelMessages("565663569132257283", 100, startID, "", "")
		if err != nil {
			log.Fatal(err.Error())
		}
		for counter, msg := range messages {
			if msg.Content == "" {
				blankLineCount++
			}else {
				fmt.Println(msg.Content)
			}
			if counter == 99 {
				startID = msg.ID
			}
		}
		if blankLineCount == 10 {
			return
		}
	}
}