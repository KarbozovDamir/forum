package data

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/KarbozovDamir/forum/internal/models"
	"github.com/satori/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	user *models.User
}

//GetUserByUsername - Get it man
func GetUserByUsername(username string) (user models.User, err error) {
	username = strings.ToLower(username)
	user = models.User{}
	err = Db.QueryRow("select * from Users where username = $1", username).Scan(&user.Mail, &user.Username, &user.Psw, &user.ID, &user.Code)
	return
}

//GetUserByMail - Get it man
func GetUserByMail(mail string) (user models.User, err error) {
	mail = strings.ToLower(mail)
	user = models.User{}
	err = Db.QueryRow("select * from Users where mail = $1", mail).Scan(&user.Mail, &user.Username, &user.Psw, &user.ID, &user.Code)
	return
}

//GetUserByID - Get it man
func GetUserByID(ID int) (user models.User, err error) {
	user = models.User{}
	err = Db.QueryRow("select * from Users where id = $1", ID).Scan(&user.Mail, &user.Username, &user.Psw, &user.ID, &user.Code)
	return
}

//CreateAndSetSession - Method for user
func (service *UserService) CreateAndSetSession(w http.ResponseWriter, r *http.Request) (err error) {
	uid := uuid.NewV4()
	hour := time.Duration(24)
	if r.Form["checkbox"] != nil {
		hour *= 30
	}
	expiration := time.Now().Add(time.Hour * hour)
	Db.Exec("delete from Session where UserId = $1", service.user.ID)
	_, err = Db.Exec("insert into Session(UserId, Uuid, Time) values ($1, $2, $3)", service.user.ID, uid, expiration)
	cookie := http.Cookie{Name: "Cookie", Value: fmt.Sprintf("%s", uid), Expires: expiration, HttpOnly: true}
	http.SetCookie(w, &cookie)
	return
}

//Update - update password
func (service *UserService) Update(w http.ResponseWriter, r *http.Request, newpass string) {
	pswEncrypted, _ := bcrypt.GenerateFromPassword([]byte(newpass), 1)
	Db.Exec("update Users set psw = $1 where username = $2", pswEncrypted, service.user.Username)
}
