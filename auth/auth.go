package auth

import (
	"seckshop/models"
	"seckshop/repo"
)

func User(token string) (user models.User, validate bool) {
	return repo.NewUserRepo().GetTokenUser(token)
}
