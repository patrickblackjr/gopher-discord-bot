package automod

import (
	"fmt"
	"os"
	"sync"

	"github.com/bwmarrin/discordgo"
)

var (
	GuildId   string
	ChannelId string
)

func AutomodInit(session *discordgo.Session) {
	GuildId = "1092598214781309029"
	ChannelId = "1092598215427240008"

	session.Identify.Intents |= discordgo.IntentAutoModerationExecution
	session.Identify.Intents |= discordgo.IntentMessageContent

	enabled := true
	rule, err := session.AutoModerationRuleCreate(GuildId, &discordgo.AutoModerationRule{
		Name:        "Automod Example",
		EventType:   discordgo.AutoModerationEventMessageSend,
		TriggerType: discordgo.AutoModerationEventTriggerKeyword,
		TriggerMetadata: &discordgo.AutoModerationTriggerMetadata{
			KeywordFilter: []string{"*cat*"},
			RegexPatterns: []string{"(c|b)at"},
		},
		Enabled: &enabled,
		Actions: []discordgo.AutoModerationAction{
			{Type: discordgo.AutoModerationRuleActionBlockMessage},
		},
	})
	if err != nil {
		panic(err)
	}

	// Print successful creation of the rule above
	fmt.Println("Successfully created automod rule.")

	// When the process ends delete the automod rule.
	defer session.AutoModerationRuleDelete(GuildId, rule.ID)

	session.AddHandlerOnce(func(se *discordgo.Session, e *discordgo.AutoModerationActionExecution) {
		_, err := session.AutoModerationRuleEdit(GuildId, rule.ID, &discordgo.AutoModerationRule{
			TriggerMetadata: &discordgo.AutoModerationTriggerMetadata{
				KeywordFilter: []string{"cat"},
			},
			Actions: []discordgo.AutoModerationAction{
				{Type: discordgo.AutoModerationRuleActionTimeout, Metadata: &discordgo.AutoModerationActionMetadata{Duration: 60}},
				{Type: discordgo.AutoModerationRuleActionSendAlertMessage, Metadata: &discordgo.AutoModerationActionMetadata{
					ChannelID: e.ChannelID,
				}},
			},
		})
		if err != nil {
			session.AutoModerationRuleDelete(GuildId, rule.ID)
		}

		se.ChannelMessageSend(e.ChannelID, "Congratulations! You have just triggered an auto moderation rule.\n"+
			"The current trigger can match anywhere in the word, so even if you write the trigger word as a part of another word, it will still match.\n"+
			"The rule has now been changed, now the trigger matches only in the full words.\n"+
			"Additionally, when you send a message, an alert will be sent to this channel and you will be **timed out** for a minute.\n")

		var counter int
		var counterMutex sync.Mutex

		session.AddHandlerOnce(func(se *discordgo.Session, e *discordgo.AutoModerationActionExecution) {
			action := "unknown"
			switch e.Action.Type {
			case discordgo.AutoModerationRuleActionBlockMessage:
				action = "block message"
			case discordgo.AutoModerationRuleActionSendAlertMessage:
				action = "send alert message into <#" + e.Action.Metadata.ChannelID + ">"
			case discordgo.AutoModerationRuleActionTimeout:
				action = "timeout"
			}

			counterMutex.Lock()
			counter++
			if counter == 1 {
				counterMutex.Unlock()
				se.ChannelMessageSend(e.ChannelID, "Nothing has changed, right? "+
					"Well, since separate gateway events are fired per each action (current is "+action+"), "+
					"you'll see a second message about an action pop up soon")
			} else if counter == 2 {
				counterMutex.Unlock()
				session.ChannelMessageSend(e.ChannelID, "Now the second ("+action+") action got executed.")
				session.ChannelMessageSend(e.ChannelID, "And... you've made it! That's the end of the example.\n"+
					"For more information about the automod and how to use it, "+
					"you can visit the official Discord docs: https://discord.dev/resources/auto-moderation or ask in our server: https://discord.gg/6dzbuDpSWY",
				)

				session.Close()
				session.AutoModerationRuleDelete(GuildId, rule.ID)
				os.Exit(0)
			}

		})
	})
}
