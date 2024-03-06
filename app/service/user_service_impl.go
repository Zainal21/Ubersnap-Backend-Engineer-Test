package service

import (
	"github.com/Zainal21/Ubersnap-backend-test/app/repositories"
)

type userServiceImpl struct {
	repo repositories.UserRepository
}

func NewUserServiceImpl(repo repositories.UserRepository) UserService {
	return &userServiceImpl{
		repo: repo,
	}
}
