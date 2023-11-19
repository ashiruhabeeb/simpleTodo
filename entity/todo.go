package entity

type Todo struct {
	TodoID		int		`json:"todo_id"`
	Title		string	`json:"title"`
	Description string	`json:"description"`
	Completed	bool	`json:"completed"`
	Date		string	`json:"date"`
	StartAt		string	`json:"start_at"`
	EndAt		string	`json:"end_at"`
}
