package repository

var (
	// PSQL queries
	insertTodo = `insert into todo (title, description, start_at, end_at) values ($1, $2, $3, $4) returning todo_id;`
)
