package userController

import (
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Get All Users"))
}
