package main

import (
	"html/template"
	"net/http"

	railscook "github.com/kalelc/go-rails-cook"
)

func Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		templates := template.Must(template.ParseFiles("views/top.html", "views/index.html"))

		if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func Result(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("cookie")
	secret := r.FormValue("secret")
	salt := r.FormValue("salt")

	cookie := railscook.Cookie{Value: value, SecretKeyBase: secret, Salt: salt}
	cookie.Decrypt()

	templates := template.Must(template.ParseFiles("views/result.html"))

	if err := templates.ExecuteTemplate(w, "result.html", cookie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
