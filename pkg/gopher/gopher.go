package gopher

import (
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/patrickblackjr/gopher-discord-bot/internal/pkg/logging"
)

var Session *discordgo.Session

func InitBot() {
	log := logging.New()
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	t := os.Getenv("BOT_TOKEN")

	Session, err = discordgo.New("Bot " + t)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
	Session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Infof("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err = Session.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	logging.New()

	defer Session.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Info("Press Ctrl+C to exit")
	// .Println("Press Ctrl+C to exit")
	<-stop

	log.Info("Gracefully shutting down.")
}
