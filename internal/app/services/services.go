package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/felipeps/user-post-comments/internal/app/cache"
	"github.com/felipeps/user-post-comments/internal/app/models"
)

func GetPosts(start, size int, ch chan models.ServiceResponse) {
	var posts []models.Post
	postsCache, exists := cache.GetPostsCache()

	if exists {
		posts = (*postsCache)
	} else {
		res, err := http.Get("https://jsonplaceholder.typicode.com/posts")

		if err != nil {
			ch <- models.ServiceResponse{
				Error: errors.New("internal server error"),
			}
		}

		defer res.Body.Close()

		posts = []models.Post{}

		if err = json.NewDecoder(res.Body).Decode(&posts); err != nil {
			ch <- models.ServiceResponse{
				Error: errors.New("internal server error"),
			}
			return
		}
	}

	cache.SetPostsCache(&posts)

	if start >= len(posts) || size > len(posts) {
		ch <- models.ServiceResponse{
			Data: &[]models.Post{},
		}
		return
	}

	posts = posts[start : start+size]

	ch <- models.ServiceResponse{
		Data: &posts,
	}
}

func GetUsers(ch chan models.ServiceResponse) {
	var users []models.User

	usersCache, exists := cache.GetUsersCache()

	if exists {
		users = *usersCache
		ch <- models.ServiceResponse{
			Data: &users,
		}
	}

	res, err := http.Get("https://jsonplaceholder.typicode.com/users")

	if err != nil {
		ch <- models.ServiceResponse{
			Error: errors.New("internal server error"),
		}
	}

	defer res.Body.Close()

	users = []models.User{}

	if err = json.NewDecoder(res.Body).Decode(&users); err != nil {
		ch <- models.ServiceResponse{
			Error: errors.New("internal server error"),
		}
	}

	cache.SetUsersCache(&users)

	ch <- models.ServiceResponse{
		Data: &users,
	}
}

func GetComments(postIds []int, ch chan models.ServiceResponse) {
	comments := make(map[int]*[]models.Comment)

	for _, postId := range postIds {
		var postComments []models.Comment

		if cacheComments, exists := cache.GetCommentsCache(postId); exists {
			comments[postId] = cacheComments
			continue
		}

		res, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d/comments", postId))

		if err != nil {
			ch <- models.ServiceResponse{
				Error: errors.New("internal server error"),
			}
		}

		defer res.Body.Close()

		postComments = []models.Comment{}

		if err = json.NewDecoder(res.Body).Decode(&postComments); err != nil {
			ch <- models.ServiceResponse{
				Error: errors.New("internal server error"),
			}
		}

		comments[postId] = &postComments

		cache.SetCommentsCache(postId, &postComments)
	}

	ch <- models.ServiceResponse{
		Data: comments,
	}
}
