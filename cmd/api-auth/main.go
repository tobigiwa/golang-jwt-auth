package main

import (
	"context"
	"log"
	"net/http"

	handlers "github.com/giwatobi/golang-jwt-auth/internal/handlers"
	"github.com/giwatobi/golang-jwt-auth/internal/store"
)

func main() {
	var ctx context.Context = context.Background()

	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/welcome", handlers.WelcomePage)
	http.HandleFunc("/signup", handlers.SignUp)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/logout", handlers.Logout)

	db_handler, err := DbSetup(ctx)
	if err != nil {
		log.Fatal(err)
	} else {

		DB_handle := store.New(db_handler)

		log.Fatal(http.ListenAndServe(":3434", nil))
	}

}
