package repository

var (
	// PSQL queries
	insertTodo = `insert into todo (title, description, start_at, end_at) values ($1, $2, $3, $4) returning todo_id;`
	getTodo = `select * from todo where todo_id = $1;`
	getTodos = `select * from todo;`
	deleteTodo = `delete from todo where todo_id = $1;`
	updateTodo = `update todo set title = $2, description = $3, start_at = $4, end_at = $5 where todo_id = $1;` 
)
