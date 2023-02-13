package Classes

type Homework struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	DueDate string `json:"duedate"` //Format: DD/MM/YYYY
	Status  bool   `json:"status"`
}
