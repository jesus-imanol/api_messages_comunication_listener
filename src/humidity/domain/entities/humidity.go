package entities

type Humidity struct {
	ID      int64 `json:"id"`
	Type string `json:"type"`
	Quantity float64 `json:"quantity"`
	Text string `json:"text"`
}