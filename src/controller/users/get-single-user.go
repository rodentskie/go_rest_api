package userController

import (
	"net/http"
)

func GetSingleUser(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Get Single User"))
}
