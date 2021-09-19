package repository

type TodoDbModel struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}
