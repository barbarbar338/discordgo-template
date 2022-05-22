package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"argon/commands"
	"argon/config"
	"argon/events"
	"argon/logger"
	"argon/utils"

	"github.com/bwmarrin/discordgo"
)

var (
	session		*discordgo.Session
)

func main() {
	var flagMigrateCommands bool

	flag.BoolVar(&flagMigrateCommands, "commands", false, "Komutları güncelle")

	flag.Parse()

	s, err := discordgo.New(fmt.Sprintf("Bot %v", config.Token))
	if err != nil {
		logger.Logger.FatalF("Discord oturumu oluşturulurken bir hata oluştu: %v", err.Error())
	}

	session = s

	s.AddHandler(events.Ready)
	s.AddHandler(events.InteractionCreate)

	err = s.Open()
	if err != nil {
		logger.Logger.FatalF("Discord API'ye bağlanılamadı: %v", err.Error())
	}

	if flagMigrateCommands {
		commands.Migrate(s)
	}
	
	shutdownChannel()
	onShutDown()
}

func shutdownChannel() {
	sc := make(chan os.Signal, 1)
	signal.Notify(
		sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		os.Interrupt,
	)

	<-sc
}

func onShutDown() {
	logger.Logger.Info("Bot durduruluyor...")

	session.Close()
	if events.StatusInterval != nil {
		utils.ClearInterval(events.StatusInterval)
	} 
}

