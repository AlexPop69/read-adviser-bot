package telegram

import (
	"net/http"
)

type Client struct {
	host     string      // хост API сервиса телеграм
	basePath string      // базовый путь, - префикс с которого начинаются все запросы
	client   http.Client // http-client
}

// функция создает клинет. Принимает хост и токен, возвращает клиент
func New(host string, token string) Client {
	return Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

// функция принимает токен и возвращает basePath
func newBasePath(token string) string {
	return "bot" + token
}
