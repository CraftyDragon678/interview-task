package cmd

import (
	"fmt"
	"xsi/framework"

	"github.com/bwmarrin/discordgo"
)

// BrCommand BR31
func BrCommand(ctx *framework.Context) {
	// game := br.NewGame()
	embed := discordgo.MessageEmbed{
		Title: "배라 게임할사람~~",
		Color: 0xcc22ff,
	}
	msg := ctx.ReplyEmbed(&embed)
	ctx.Discord.MessageReactionAdd(ctx.TextChannel.ID, msg.ID, "👍")
	ctx.Discord.MessageReactionAdd(ctx.TextChannel.ID, msg.ID, "👎")

	userID, react := framework.NextReactionAddForAll(ctx.Discord, ctx.TextChannel.ID, msg.ID, []string{"👍", "👎"}, 10)
	fmt.Println(userID, react)
}
