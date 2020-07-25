package routes

import (
	"Webtutorial/session"
	"net/http"
)

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		session, _ := session.Store.Get(r, "session")
		delete(session.Values, "user_id")
		session.Save(r, w)

		http.Redirect(w, r, "/login", 302)
	}
}
