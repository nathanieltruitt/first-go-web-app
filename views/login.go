package views

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Ninja struct {
	Name string
}

var userDB = map[string]string{
	"derpy@yes.com": "ickpass",
}

func Login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./views/html/login.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, Ninja{Name: "Please login"})
}

func LoginSubmit(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	if userDB[email] == password {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "You have logged in successfully.")
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "You have failed to login.")
	}
}