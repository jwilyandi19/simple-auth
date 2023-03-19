package repository

const (
	insertUser  = `INSERT users SET user_name=?, user_full_name=?, user_password=?`
	getUser     = `SELECT user_id, user_name, user_full_name, user_password FROM users WHERE user_name = ?`
	fetchUsers  = `SELECT user_id, user_name, user_full_name FROM users LIMIT ?, ?`
	isUserExist = `SELECT user_name FROM users WHERE user_name = ?`
)
