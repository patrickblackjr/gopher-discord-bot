package callbacks

import "github.com/bwmarrin/discordgo"

const (
	voiceStateUpdate           = "VoiceStateUpdate"
	voiceStateUpdateEventError = "Unable to process event: " + voiceStateUpdate
)

type voiceStateUpdateMetadata struct {
	Session       *discordgo.Session
	Guild         *discordgo.Guild
	Member        *discordgo.Member
	Channel       *discordgo.Channel
	EphemeralRole *discordgo.Role
}

func (handler *Handler) VoiceStateUpdate(session *discordgo.Session, voiceState *discordgo.VoiceStateUpdate) {

}
