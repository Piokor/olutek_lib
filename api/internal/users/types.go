package users

type User struct {
	Id       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}
