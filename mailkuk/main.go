package main

import (
	"github.com/charmbracelet/log"
	"os"
	"os/signal"
	"syscall"
)

const VERSION = "23.0"

func main() {
	log.Infof("mailkuk version %s", VERSION)

	config, err := loadConfig()
	if err != nil {
		log.Error("Error while loading config file", "err", err)
		os.Exit(1)
	}

	if config.Server.Debug {
		log.SetLevel(log.DebugLevel)
	}

	loadRouter(config.Routing)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		if SMTPServer != nil {
			SMTPServer.Close()
		}
	}()

	log.Info("Starting SMTP server", "addr", config.Server.ListenAddr)
	err = startServer(config.Server)
	if err != nil {
		log.Error("Error while running SMTP server", "err", err)
		os.Exit(1)
	}
}
