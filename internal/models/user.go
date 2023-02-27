package models

//User - special struct for db
type User struct {
	Mail     string
	Username string
	Psw      string
	ID       int
	Code     string
}
