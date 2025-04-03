package entities

type MessageFertilizer struct {
  Type string `json:"type"`
  Command string `json:"command"`
  Quantity float64 `json:"quantity"`
  User string `json:"user"`
}