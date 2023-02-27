package router

import (
	"net/http"

	data "github.com/KarbozovDamir/forum/data"
)

//LogOut - log out
func LogOut(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	cookie, err := r.Cookie("Cookie")
	if err != http.ErrNoCookie {
		UUID := cookie.Value
		data.DeleteByUUID(UUID)
		cookie.MaxAge = -1 //equivalently Max-Age: 0 with delete
		http.SetCookie(w, cookie)
	}
	http.Redirect(w, r, "/", code)
	Parser.Authorised = false
	Parser.ID = 0
}
