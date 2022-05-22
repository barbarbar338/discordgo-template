package commands

import (
	"argon/logger"

	"github.com/bwmarrin/discordgo"
)

var (
	commands = []*discordgo.ApplicationCommand{
		&PingData,
	}
	
	Handlers = map[string]func(*discordgo.Session, *discordgo.InteractionCreate) {
		"ping": PingHandler,
	}
)

func Migrate(s *discordgo.Session) {
	logger.Logger.Info("Komutlar yükleniyor...")
	for _, command := range commands {
		logger.Logger.InfoF("Yüklenen komut: %v", command.Name)
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", command)
		if err != nil {
			logger.Logger.WarningF("Komut yüklenirken hata oluştu: %s", err.Error())
		}
	}
}
