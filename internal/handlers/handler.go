package handlers

import "net/http"

func HomePage(w http.ResponseWriter, req *http.Request) {
	w.Header()
	w.Write([]byte("This page is available to every one."))

}

func SignUp(w http.ResponseWriter, req *http.Request) {

}

func WelcomePage(w http.ResponseWriter, req *http.Request) {

}

func Login(w http.ResponseWriter, req *http.Request) {

}
func Logout(w http.ResponseWriter, req *http.Request) {

}
