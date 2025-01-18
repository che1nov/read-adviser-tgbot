package main

import (
	"context"
	"flag"
	tgClient "github.com/che1nov/read-adviser-tgbot/clients/telegram"
	event_consumer "github.com/che1nov/read-adviser-tgbot/consumer/event-consumer"
	"github.com/che1nov/read-adviser-tgbot/events/telegram"
	"github.com/che1nov/read-adviser-tgbot/storage/sqlite"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const (
	tgBotHost         = "api.telegram.org"
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize         = 100
)

// 7478806356:AAHocGugABm4psU77XFOmlrca9xcPNk7fy8
func main() {
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatalf("can't connect to storage", err)
	}

	s.Init(context.TODO())

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		s,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
