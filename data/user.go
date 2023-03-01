package data

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/satori/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string
	MaxAge     int
	Secure     bool
	HTTPOnly   bool
	Raw        string
	Unparsed   []string // Raw text of unparsed attribute-value pairs
}

// User - special struct for db
type User struct {
	Mail     string
	Username string
	Psw      string
	ID       int
	Code     string
}

// GetUserByUsername - Get it man
func GetUserByUsername(username string) (user User, err error) {
	username = strings.ToLower(username)
	user = User{}
	err = Db.QueryRow("select * from Users where username = $1", username).Scan(
		&user.Mail,
		&user.Username,
		&user.Psw,
		&user.ID,
		&user.Code)
	return
}

// GetUserByMail - Get it man
func GetUserByMail(mail string) (user User, err error) {
	mail = strings.ToLower(mail)
	user = User{}
	err = Db.QueryRow("select * from Users where mail = $1", mail).Scan(
		&user.Mail,
		&user.Username,
		&user.Psw,
		&user.ID,
		&user.Code)
	return
}

// GetUserByID - Get it man
func GetUserByID(ID int) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select * from Users where id = $1", ID).Scan(&user.Mail, &user.Username, &user.Psw, &user.ID, &user.Code)
	return
}

// CreateAndSetSession - Method for user
func (service *User) CreateAndSetSession(w http.ResponseWriter, r *http.Request) (err error) {
	uid := uuid.NewV4()
	hour := time.Duration(24)
	if r.Form["checkbox"] != nil {
		hour *= 30
	}
	expiration := time.Now().Add(time.Hour * hour)
	Db.Exec("delete from Session where UserId = $1", service.ID)
	_, err = Db.Exec("insert into Session(UserId, Uuid, Time) values ($1, $2, $3)", service.ID, uid, expiration)
	cookie := http.Cookie{Name: "Cookie", Value: fmt.Sprintf("%s", uid), Expires: expiration, HttpOnly: true}
	http.SetCookie(w, &cookie)
	return
}

// Update - update password
func (service *User) Update(w http.ResponseWriter, r *http.Request, newpass string) {
	pswEncrypted, _ := bcrypt.GenerateFromPassword([]byte(newpass), 1)
	Db.Exec("update Users set psw = $1 where username = $2", pswEncrypted, service.Username)
}
