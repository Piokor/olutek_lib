package users

type User struct {
	Id       int64  `databse:"id"`
	Username string `database:"username"`
}
