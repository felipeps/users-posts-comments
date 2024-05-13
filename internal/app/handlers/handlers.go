package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/felipeps/user-post-comments/internal/app/models"
	"github.com/felipeps/user-post-comments/internal/app/services"
	"github.com/felipeps/user-post-comments/internal/app/utils"
)

func GetUsersPosts(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	start, err := strconv.Atoi(queryParams.Get("start"))

	if err != nil || start < 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"message\": \"invalid start parameter\"}"))
		return
	}

	size, err := strconv.Atoi(queryParams.Get("size"))

	if err != nil || size < 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"message\": \"invalid size parameter\"}"))
		return
	}

	uc := make(chan models.ServiceResponse)
	pc := make(chan models.ServiceResponse)

	go services.GetUsers(uc)
	go services.GetPosts(start, size, pc)

	getUsersResponse := <-uc
	getPostsResponse := <-pc

	if getUsersResponse.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\": \"internal server error\"}"))
		return
	}

	if getPostsResponse.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\": \"internal server error\"}"))
		return
	}

	users := *(getUsersResponse.Data.(*[]models.User))
	posts := *(getPostsResponse.Data.(*[]models.Post))

	if len(posts) <= 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	cc := make(chan models.ServiceResponse)

	go services.GetComments(utils.GetPostIds(&posts), cc)

	getCommentsResponse := <-cc
	comments := getCommentsResponse.Data.(map[int]*[]models.Comment)

	response := utils.Combine(&users, &posts, comments)
	out, err := json.Marshal(response)

	if err != nil {
		w.Write([]byte("internal server error"))
		return
	}

	w.Write(out)
}
