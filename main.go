package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/eoschaos/aniki-bot/commands/anime34"
	"github.com/eoschaos/aniki-bot/commands/gelbooru"
	"github.com/eoschaos/aniki-bot/commands/r34xyz"
	"github.com/eoschaos/aniki-bot/commands/rule34"
	"github.com/eoschaos/aniki-bot/utils"
	"golang.org/x/exp/slices"
)

var (
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken       = flag.String("token", "", "Bot access token")
	DeniedIDs      = flag.String("denied", "", "IDs of users who are denied access to the bot")
	SuperUserIDs   = flag.String("superusers", "", "IDs of users who are super users for the bot")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutting down or not")
)

var s *discordgo.Session

func init() { flag.Parse() }

func init() {
	var err error
	s, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

var (
	commands = []*discordgo.ApplicationCommand{
		animer34.Command,
		gelbooru.Command("gelbooru"),
		gelbooru.Command("gel"),
		r34xyz.Command,
		rule34.Command,
	}

	commandCoolDowns = map[string]map[string]time.Time{}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ar34":     animer34.Handler(),
		"gelbooru": gelbooru.Handler(),
		"gel":      gelbooru.Handler(),
		"r34":      rule34.Handler(),
		"xyz34":    r34xyz.Handler(),
	}
)

func init() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		commandName := i.ApplicationCommandData().Name
		if h, ok := commandHandlers[commandName]; ok {
			if i.Type == discordgo.InteractionApplicationCommandAutocomplete {
				h(s, i)
				return
			}

			channel, _ := s.Channel(i.ChannelID)
			if !channel.NSFW {
				utils.SendEmbedArticle(s, i, "Go be horny in the NSFW channels!")
				return
			}

			if slices.Contains(strings.Split(*DeniedIDs, ","), i.Member.User.ID) {
				utils.SendEmbedArticle(s, i, fmt.Sprintf("You're denied access to the bot, %s", i.Member.User.Username))
				return
			}

			coolDown := commandCoolDowns[commandName]
			lastReply := coolDown[i.Member.User.ID]
			if lastReply.IsZero() {
				h(s, i)
				coolDown[i.Member.User.ID] = time.Now()
				return
			}

			current := time.Now()
			isSuperUser := slices.Contains(strings.Split(*SuperUserIDs, ","), i.Member.User.ID)
			if current.After(lastReply.Add(5*time.Second)) || isSuperUser {
				h(s, i)
				coolDown[i.Member.User.ID] = time.Now()
				return
			}

			utils.SendEmbedArticle(s, i, "Calm down, don't choke the turkey!")
		}
	})
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, *GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		commandCoolDowns[v.Name] = map[string]time.Time{}
		registeredCommands[i] = cmd
	}

	defer func(s *discordgo.Session) {
		_ = s.Close()
	}(s)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	if *RemoveCommands {
		log.Println("Removing commands...")
		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, *GuildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	log.Println("Gracefully shutting down.")
}
