package daos

import "github.com/AndrewJTo/htmx-forum/models"

func CreatePost(newPost models.Post) (models.Post, error) {
	posts = append(posts, newPost)

	return newPost, nil
}
