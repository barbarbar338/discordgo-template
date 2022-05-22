package config

import (
	"os"

	"argon/logger"

	"github.com/joho/godotenv"
)

var (
	Token			string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Logger.Fatal(".env dosyası yüklenirken hata oluştu: " + err.Error())
	}

	Token = os.Getenv("BOT_TOKEN")
}
