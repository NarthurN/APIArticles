package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/NarthurN/APIArticles/pkg/mocks"
	"github.com/gorilla/mux"
)

func GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, article := range mocks.Articles {
		if article.Id == id {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(article)
			break
		}
	}
}
