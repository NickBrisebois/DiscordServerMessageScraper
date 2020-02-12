package biglez_scraper

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type ServerScraper struct {
	botConf *Config
}

func NewServerScraper(config *Config) *ServerScraper {
	return &ServerScraper{
		botConf: config,
	}
}

func (sc *ServerScraper) InitScraper() error {
	log.Println("Initializing The Big LezTM Server Scraper")

	discord, err := discordgo.New("Bot " + sc.botConf.DiscordToken)

	if err != nil {
		return err
	}

	discord.AddHandler(messageCreate)

	err = discord.Open()
	if err != nil {
		return err
	}

	sigCapture := make(chan os.Signal, 1)
	signal.Notify(sigCapture, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sigCapture

	discord.Close()
	return nil
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "who ARE you?" {
		s.ChannelMessageSend(m.ChannelID, "Have no fear, I am a fellow human just like you")
	}
}