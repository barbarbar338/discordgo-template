package events

import (
	"time"

	"argon/logger"
	"argon/utils"

	"github.com/bwmarrin/discordgo"
)

const (
	statusIntervalPeriot	= 30 * time.Second
)

var (
	statusTexts 			= []string {
		"💩",
		"👀",
	}
	StatusInterval			chan bool
)

func Ready(session *discordgo.Session, event *discordgo.Ready) {	
	err := session.UpdateGameStatus(0, utils.RandomString(statusTexts))
	if err != nil {
		logger.Logger.WarningF("Oynuyor mesajı güncellenirken bir hata oluştu: %s", err.Error())
	}

	StatusInterval = utils.SetInterval(func() {
		err := session.UpdateGameStatus(0, utils.RandomString(statusTexts))
		if err != nil {
			logger.Logger.WarningF("Oynuyor mesajı güncellenirken bir hata oluştu: %s", err.Error())
		}
	}, statusIntervalPeriot)

	logger.Logger.InfoF("[%s:%s] Hazır!", session.State.User.Username, session.State.User.ID)
}
