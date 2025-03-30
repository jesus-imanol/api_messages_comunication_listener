package entities

type Message struct {
	ID      int64 `json:"id"`
	Type string `json:"type"`
	Quantity float64 `json:"quantity"`
	Text string `json:"text"`
	User string `json:"username"`
	CreatedAt string `json:"created_at"`
}