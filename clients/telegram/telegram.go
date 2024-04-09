package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host     string      // хост API сервиса телеграм
	basePath string      // базовый путь, - префикс с которого начинаются все запросы
	client   http.Client // http-client
}

const (
	getUpdatesMethod  = "getUpdates"
	sendmessageMethod = "sendMessage"
)

// the function creates the client. Accepts host and token, returns client
func New(host string, token string) Client {
	return Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

// function takes a token and returns basePath
func newBasePath(token string) string {
	return "bot" + token
}

// Method for getting updates
// https://core.telegram.org/bots/api#getting-updates
func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	// adding parameters offset and limit to request
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, err
	}

	var res UpdateResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}

// Use this method to send text messages. On success, the sent Message is returned.
// https://core.telegram.org/bots/api#sendmessage
func (c *Client) SendMessage(chatID int, text string) error {
	// adding parameters offset and limit to request
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequest(sendmessageMethod, q)
	if err != nil {
		return fmt.Errorf("can't send message: %w", err)
	}

	return nil
}

// method for sending a request. Takes method and query, returns data and error
func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("can`t do request: %w", err)
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can`t do request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can`t do request: %w", err)
	}

	return body, nil
}
