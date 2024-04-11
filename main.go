package main

import (
	"flag"
	"log"

	tgClient "read-adviser-bot/clients/telegram"
	eventconsumer "read-adviser-bot/consumer/event-consumer"
	"read-adviser-bot/events/telegram"
	"read-adviser-bot/storage/files"
)

const (
	storagePath = "file-storage"
	batchSize   = 10
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.New("api.telegram.org", mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal()
	}
}

// функция для получения токена
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

// func mustHost() string {
// 	host := flag.String(
// 		"tg-bot-host",
// 		"api.telegram.org",
// 		"api.telegram.org",
// 	)

// 	flag.Parse()

// 	if *host == "" {
// 		log.Fatal("host is not specified")
// 	}

// 	return *host
// }
