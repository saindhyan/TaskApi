package entity

type Task struct {
	ID          int    `json:"Id"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	DueDate     string `json:"Due Date"`
	Status      string `json:"Status"`
}
