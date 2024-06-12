package daos

import (
	"errors"

	"github.com/AndrewJTo/htmx-forum/models"
)

func FindUserByEmail(email string) (models.User, error) {
	for _, u := range users {
		if u.Email == email {
			return u, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func CreateUser(newUser models.User) (models.User, error) {
	users = append(users, newUser)
	return newUser, nil
}
