package utils

import (
	"fmt"
	"strings"

	data "github.com/KarbozovDamir/forum/internal/data"
	"golang.org/x/crypto/bcrypt"
)

// AddUser - allow us add user to the database
func AddUser(username, mail, psw string) {
	pswEncrypted, _ := bcrypt.GenerateFromPassword([]byte(psw), 1)

	code := String(6)
	_, err3 := data.Db.Exec("insert into Users(mail, username, psw, code) values ($1, $2, $3, $4)", strings.ToLower(mail), strings.ToLower(username), pswEncrypted, code)
	SendRestoreCodeToUser(mail, code)
	if err3 != nil {
		fmt.Println("AddUser is broken 3")
		fmt.Println(err3)
	}
}
