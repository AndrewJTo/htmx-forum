package daos

import (
	"errors"

	"github.com/AndrewJTo/htmx-forum/models"
)

func FindThreadById(threadId int) (models.Thread, error) {
	for _, t := range threads {
		if t.Id == threadId {
			return t, nil
		}
	}
	return models.Thread{}, errors.New("category not found")
}

func CreateThread(newThread models.Thread) (models.Thread, error) {
	threads = append(threads, newThread)

	return newThread, nil
}

func FindThreadPosts(thread *models.Thread) ([]models.Post, error) {
	var postsList []models.Post
	for _, p := range posts {
		if p.ThreadId == thread.Id {
			postsList = append(postsList, p)
		}
	}
	return postsList, nil
}
