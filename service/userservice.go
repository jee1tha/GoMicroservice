package service

import (
	"GoMicroservice/database"
	"GoMicroservice/entities"
)

func CreateUserService(user entities.User) {
	database.Instance.Create(user)
}
