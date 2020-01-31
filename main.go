package main

import (
	"gopractice/config"
	"gopractice/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println(config.Config.LogFile)
	log.Println(config.Config.Port)
	log.Println(config.Config.Name)
	log.Println(config.Config.Driver)
}
