package entities

type Message struct {
	ID      int64 `json:"id"`
	Type string `json:"type"`
	Quantity float64 `json:"quantity"`
	Text string `json:"text"`
}
func NewMessage(typing string, quantity float64, text string) *Message {
	return &Message{
		Type: typing,
		Quantity: quantity,
		Text: text,
	}
}