package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"xsi/cmd"
	"xsi/framework"

	"github.com/bwmarrin/discordgo"

	"github.com/joho/godotenv"
)

var (
	bots       [1]*discordgo.Session
	prefix     string = "!"
	botID      string
	cmdHandler *framework.CommandHandler
)

func runBot(token string, shardID int) *discordgo.Session {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalln("Making bot\n", err)
	}

	discord.ShardCount = len(bots)
	discord.ShardID = shardID
	discord.State.MaxMessageCount = 1000

	discord.AddHandler(commandHandler)

	err = discord.Open()
	if err != nil {
		log.Fatalln("Opening websocket\n", err)
	}

	return discord
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file. Please Make sure your .env exist")
	}
}

func main() {
	token := os.Getenv("TOKEN")
	if _prefix := os.Getenv("prefix"); _prefix != "" {
		prefix = _prefix
	}

	framework.InitDB()
	defer framework.ExitDB()

	cmdHandler = framework.NewCommandHandler()
	registerCommands()
	framework.CmdHandler = cmdHandler

	for i := 0; i < len(bots); i++ {
		bots[i] = runBot(token, i)
		defer bots[i].Close()
	}

	usr, err := bots[0].User("@me")
	if err != nil {
		log.Fatalln("obtaining account datils\n", err)
	}

	botID = usr.ID

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func registerCommands() {
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == botID || user.Bot {
		return
	}
	content := message.Content
	if len(content) <= len(prefix) {
		return
	}
	if content[:len(prefix)] != prefix {
		return
	}
	content = content[len(prefix):]
	if len(content) < 1 {
		return
	}

	if !framework.CheckUserExist(user.ID) {
		framework.AddUser(user.ID)
	}

	args := strings.Fields(content)
	name := strings.ToLower(args[0])
	command, found := cmdHandler.Get(name)
	if !found {
		return
	}
	channel, err := discord.State.Channel(message.ChannelID)
	if err != nil {
		fmt.Println("[ERROR] getting Channel", err)
	}

	var guild *discordgo.Guild
	if channel.Type != discordgo.ChannelTypeDM {
		guild, err = discord.State.Guild(channel.GuildID)
		if err != nil {
			fmt.Println("[ERROR] getting Guild", err)
		}
	}

	ctx := framework.NewContext(discord, guild, channel, user, message.Message)
	ctx.Args = args[1:]
	c := *command
	go c(ctx)
}
