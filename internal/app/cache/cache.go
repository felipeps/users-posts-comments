package cache

import "github.com/felipeps/user-post-comments/internal/app/models"

var postsCache = &[]models.Post{}
var usersCache = &[]models.User{}
var commentsCache = map[int]*[]models.Comment{}

func GetPostsCache() (*[]models.Post, bool) {
	return postsCache, len(*postsCache) > 0
}

func SetPostsCache(posts *[]models.Post) {
	postsCache = posts
}

func GetUsersCache() (*[]models.User, bool) {
	return usersCache, len(*usersCache) > 0
}

func SetUsersCache(users *[]models.User) {
	usersCache = users
}

func GetCommentsCache(postId int) (*[]models.Comment, bool) {
	comments, ok := commentsCache[postId]

	if !ok {
		return nil, false
	}

	return comments, true
}

func SetCommentsCache(postId int, comments *[]models.Comment) {
	value, ok := commentsCache[postId]

	if !ok {
		commentsCache[postId] = comments
	} else {
		*value = append(*value, *comments...)
	}
}
