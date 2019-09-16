package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Beetrack/rails-cookie-go/decrypt"
)

func Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		templates := template.Must(template.ParseFiles("views/index.html"))

		if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func Result(w http.ResponseWriter, r *http.Request) {
	cookie := r.FormValue("name")
	salt := r.FormValue("salt")
	secret := r.FormValue("secret")
	cookie := decrypt.Rails5Cookie{Value: cookie, SecretKeyBase: secret, Salt: salt}
	fmt.Println(cookie)

}
