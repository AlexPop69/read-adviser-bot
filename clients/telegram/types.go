package telegram

// The response contains a JSON object, which always has a Boolean field 'ok'.
// If 'ok' equals True, the request was successful and the result of the query can be found in the 'result' field.
// https://core.telegram.org/bots/api#making-requests
type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"` // slice of request result
}

// This object represents an incoming update.
// https://core.telegram.org/bots/api#getting-updates
type Update struct {
	ID      int              `json:"update_id"` // The update's unique identifier
	Message *IncomingMessage `json:"message"`   // New incoming message
}

// This object represents an incoming message. https://core.telegram.org/bots/api#message
type IncomingMessage struct {
	Text string `json:"text"` // For text messages, the actual UTF-8 text of the message
	From From   `json:"from"` // Sender of the message
	Chat Chat   `json:"chat"` // Chat the message belongs to

}

type From struct {
	Username string `json:"username"` // User's or bot's username
}

type Chat struct {
	ID int `json:"id"` // Unique identifier for this chat
}

type Message struct {
	ID   int    //
	Text string //
}
