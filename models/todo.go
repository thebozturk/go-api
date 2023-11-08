package models

// Todo is a struct for todo
type Todo struct {
	Title   string `json:"title.omitempty"`
	Content string `json:"content,omitempty"`
}
