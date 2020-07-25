package routes

import (
	"Webtutorial/middleware"
	"Webtutorial/models"
	"Webtutorial/session"
	"Webtutorial/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", middleware.AuthRequired(indexHandler))
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/logout", logoutHandler)
	r.HandleFunc("/register", registerHandler)
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	//	r.HandleFunc("/{username}",middleware.AuthRequired(userGetHandler))
	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		updates, err := models.GetUpdate()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}
		utils.ExecuteTemplate(w, "index.html", updates)
	}

	if r.Method == "POST" {
		session, _ := session.Store.Get(r, "session")
		untypedUserId := session.Values["user_id"]
		userId, ok := untypedUserId.(int64)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}
		r.ParseForm()
		body := r.PostForm.Get("update")
		err := models.PostUpdate(userId, body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}
		http.Redirect(w, r, "/", 302)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	//ctx := context.TODO()
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, "login.html", nil)
	}

	if r.Method == "POST" {
		r.ParseForm()
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")
		user, err := models.AuthenticateUser(username, password)
		if err != nil {
			switch err {
			case models.ErrUserNotFound:
				utils.ExecuteTemplate(w, "login.html", "unknown user")
			case models.ErrInvalidLoginDetails:
				utils.ExecuteTemplate(w, "login.html", "Invalid Login Details")
			default:
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
			}
			return
		}
		userId, err := user.GetId()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error "))
			return
		}
		session, _ := session.Store.Get(r, "session")
		session.Values["user_id"] = userId
		session.Save(r, w)
		http.Redirect(w, r, "/", 302)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	//ctx := context.TODO()
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, "register.html", nil)
	}

	if r.Method == "POST" {
		r.ParseForm()
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")
		err := models.RegisterUser(username, password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Sever Error"))
			return
		}

		http.Redirect(w, r, "/login", 302)

	}
}

// func userGetHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	if r.Method == "GET" {
// 		username := vars["username"]

// 	}

// }
