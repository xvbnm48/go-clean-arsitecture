package routes

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/xvbnm48/go-clean-arsitecture/entity"
	"github.com/xvbnm48/go-clean-arsitecture/repository"
)

// type Post struct {
// 	Id    int    `json:"id"`
// 	Title string `json:"title"`
// 	Text  string `json:"text"`
// }

var (
	repo repository.PostRepository = repository.NewRepository()
)

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	post.Id = rand.Int63()
	repo.Save(&post)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
