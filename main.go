package main

import (
	"flag"
	"log"

	"read-adviser-bot/clients/telegram"
)

func main() {

	tgClient := telegram.New(mustHost(), mustToken())

	// fetcher = fetcher.New(tgClient)

	// processor = processor.New(tgClient)

	// consumer.Start(fetcher, processor)
}

// функция для получения токена
func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}

func mustHost() string {
	host := flag.String(
		"host",
		"",
		"api.telegram.org",
	)

	flag.Parse()

	return *host
}
