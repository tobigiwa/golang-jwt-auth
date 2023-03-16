package middlewares

import (
	"encoding/json"
	"net/http"

	types "github.com/giwatobi/golang-jwt-auth/pkg/types"
)

func AuthenticateUser(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var Credentials types.UserCredentials

		err := json.NewDecoder(r.Body).Decode(&Credentials)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("User Credentials could not be processed"))
			return
		}

	})

}
