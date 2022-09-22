package userController

import (
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Create User"))
}
