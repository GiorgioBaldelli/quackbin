package models

type Paste struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	IsPrivate bool   `json:"is_private"`
	Password  string `json:"password,omitempty"`
}
