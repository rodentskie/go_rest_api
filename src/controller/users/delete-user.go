package userController

import (
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
