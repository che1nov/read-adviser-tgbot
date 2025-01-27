package main

import (
	"context"
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
	tgBotToken        = "7752693039:AAGEMKLZYj7mm-z8TcFJ7rOhC0AIfhdsU6s" // Жестко закодированный токен
)

func main() {
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatalf("can't connect to storage: %v", err)
	}

	s.Init(context.TODO())

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, tgBotToken),
		s,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}
