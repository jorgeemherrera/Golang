package server

import (
	"fmt"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/", index)

	http.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetCountries(w, r)

		case http.MethodPost:
			addCountry(w, r)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}
	})
}
