package utils

import "github.com/felipeps/user-post-comments/internal/app/models"

func Combine(users *[]models.User, posts *[]models.Post, comments map[int]*[]models.Comment) *[]models.UserPostComment {
	userPostComments := []models.UserPostComment{}
	userMap := make(map[int]models.User)

	for _, user := range *users {
		userMap[int(user.Id)] = user
	}

	for _, post := range *posts {
		userPostComments = append(userPostComments, models.UserPostComment{
			User:     userMap[int(post.UserId)],
			Post:     post,
			Comments: *comments[int(post.Id)],
		})
	}

	return &userPostComments
}

func GetPostIds(posts *[]models.Post) []int {
	postIds := []int{}

	for _, post := range *posts {
		postIds = append(postIds, int(post.Id))
	}

	return postIds
}
