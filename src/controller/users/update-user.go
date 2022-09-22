package userController

import (
	"net/http"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
