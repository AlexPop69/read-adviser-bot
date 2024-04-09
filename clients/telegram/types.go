package telegram

// This object represents an incoming update.
// https://core.telegram.org/bots/api#getting-updates
type Update struct {
	ID      int    `json:"update_id"` // The update's unique identifier
	Message string `json:"message"`   // New incoming message
}

// The response contains a JSON object, which always has a Boolean field 'ok'.
// If 'ok' equals True, the request was successful and the result of the query can be found in the 'result' field.
// https://core.telegram.org/bots/api#making-requests
type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"` // slice of request result
}

type Message struct {
	ID   int    //
	Text string //
}
