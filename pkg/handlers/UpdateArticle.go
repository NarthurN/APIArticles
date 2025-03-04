package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/NarthurN/APIArticles/pkg/mocks"
	"github.com/NarthurN/APIArticles/pkg/models"
	"github.com/gorilla/mux"
)

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var updatedArticle models.Article
	if err := json.Unmarshal(body, &updatedArticle); err != nil {
		log.Printf("cannot unmarshal body in UpdatedArtcicle %v", err)
	}

	for index, article := range mocks.Articles {
		if article.Id == id {
			article.Title = updatedArticle.Title
			article.Desc = updatedArticle.Desc
			article.Content = updatedArticle.Content

			mocks.Articles[index] = article

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}
}
