package framework

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func nextReactionAdd(s *discordgo.Session) *discordgo.MessageReactionAdd {
	out := make(chan *discordgo.MessageReactionAdd)
	s.AddHandlerOnce(func(_ *discordgo.Session, e *discordgo.MessageReactionAdd) {
		out <- e
	})
	return <-out
}

// NextReactionAdd wait for reaction add
func NextReactionAdd(s *discordgo.Session, channelID, userID, msgID string, emojiID []string, duration int) string {
	out := make(chan string)
	ret := "-"
	go func() {
		time.Sleep(time.Second * time.Duration(duration))
		out <- ""
	}()

	go func() {
		for {
			react := nextReactionAdd(s)
			if react.ChannelID == channelID && react.UserID == userID && react.MessageID == msgID {
				for _, emoji := range emojiID {
					if react.Emoji.MessageFormat() == emoji {
						out <- react.Emoji.MessageFormat()
						return
					}
				}
			}
			if ret == "" {
				return
			}
		}
	}()
	ret = <-out
	return ret
}

// NextReactionAddForAll wait for reaction add, all
func NextReactionAddForAll(s *discordgo.Session, channelID, msgID string, emojiID []string, duration int) (userID string, emoji string) {
	out := make(chan string)
	ret1 := "-"
	ret2 := "-"
	go func() {
		time.Sleep(time.Second * time.Duration(duration))
		out <- ""
	}()

	go func() {
		for {
			react := nextReactionAdd(s)
			if react.ChannelID == channelID && react.MessageID == msgID && !react.Emoji.User.Bot {
				for _, emoji := range emojiID {
					if react.Emoji.MessageFormat() == emoji {
						ret1 = react.UserID
						out <- react.Emoji.MessageFormat()
						return
					}
				}
			}
			if ret2 == "" {
				return
			}
		}
	}()
	ret2 = <-out
	return ret1, ret2
}
