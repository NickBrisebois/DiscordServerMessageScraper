## DiscordServerMessageScraper

This bot, when setup, will scrape every message ever posted to the server it has been added to.

All logs will be dumped into a folder ./dump/ with the name {channel-name}-dump.txt. 

### Building

To build you need these go packages:

    github.com/BurntSushi/toml
    github.com/bwmarrin/discordgo
    
To build all you need to do is CD into the folder and run:

    go build
    
### Running

To run you need to have a config.toml file with your bot API token inside. An example has been provided, just change discordToken to your token.

The bot will search for a file called config.toml by default but you can specify a specific path with -c /path/to/config.toml
    