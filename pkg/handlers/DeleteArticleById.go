package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteArticleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

}
