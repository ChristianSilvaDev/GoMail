package entity

type Mail struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	Destination string `json:"destination"`
	Subject     string `json:"subject"`
	Body        string `json:"body"`
	Status      string `json:"status"`
}
