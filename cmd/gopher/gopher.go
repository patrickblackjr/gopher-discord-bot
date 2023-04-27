package main

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentGuildMembers | discordgo.IntentGuildPresences
)

type environmentVariables struct {
	BotToken            string `env:"BOT_TOKEN,required"`
	LogLevel            string `env:"LOG_LEVEL" envDefault:"info"`
	LogTimezoneLocation string `env:"LOG_TIMEZONE_LOCATION" envDefault:"UTC"`
	BotName             string `env:"BOT_NAME" envDefault:"Ephemeral Roles"`
	RolePrefix          string `env:"ROLE_PREFIX" envDefault:"{eph}"`
	RoleColor           int    `env:"ROLE_COLOR_HEX2DEC" envDefault:"16753920"`
	InstanceName        string `env:"INSTANCE_NAME" envDefault:"ephemeral-roles-0"`
	ShardCount          int    `env:"SHARD_COUNT" envDefault:"1"`
	shardID             int
}

func (envVars *environmentVariables) parseShareID() error {
	shardIDRegEx := regexp.MustCompile(`-\d.*$`)

	shardIDString := shardIDRegEx.FindString(envVars.InstanceName)
	shardIDString = strings.TrimPrefix(shardIDString, "-")

	shardID, err := strconv.Atoi(shardIDString)
	if err != nil {
		return fmt.Errorf("error parsing share ID: %w", err)
	}

	envVars.shardID = shardID
	return nil
}

func startSession(ctx context.Context) {

}
