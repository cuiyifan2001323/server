package service

import (
	"context"
	"gin-study/domain"
	"gin-study/repository"
)

type UserService struct {
}

func (s UserService) Signup(ctx context.Context, u domain.User) error {
	repo := new(repository.UserRepository)
	return repo.Create(ctx, u)
}
