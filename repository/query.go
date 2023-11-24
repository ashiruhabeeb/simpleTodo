package repository

var (
	// psql TODO table queries
	insertTodo = `insert into todo (title, description, start_at, end_at) values ($1, $2, $3, $4) returning todo_id;`
	getTodo = `select * from todo where todo_id = $1;`
	getTodos = `select * from todo;`
	deleteTodo = `delete from todo where todo_id = $1;`
	updateTodo = `update todo set title = $2, description = $3, start_at = $4, end_at = $5 where todo_id = $1;` 

	// psql USERS table queries
	insertuser = `insert into users (username, fullname, email, password, phone, address, avatar, dob) values ($1, $2, $3, $4, $5, $6, $7) returning user_id;`
	getUser = `select todo_id, username, fullname, email, phone, address, avatar, dob, created_at, updated_at from users where user_id = $1;`
	getUserByEmail = `select * from users where email = $1;`
	getUsers = `select todo_id, username, fullname, email, phone, address, avatar, dob, created_at, updated_at from users order by user_id limit 3;`
	updateUser = `update users set username = $2, fullname = $3, phone = $4, address = $5, dob = $6 where user_id = $1;`
	deleteUser = `delete from users where user_id = $1;`
)
