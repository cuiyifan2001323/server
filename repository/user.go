package repository

import (
	"context"
	"gin-study/domain"
	"gin-study/repository/dao"
	"gorm.io/gorm"
)

type UserRepository struct {
	dao *dao.UserDAO
}

var (
	EmailConflictErr = dao.EmailConflictErr
	ErrUserNotFound  = gorm.ErrRecordNotFound
)

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}
func (repo *UserRepository) Create(ctx context.Context, u domain.User) error {
	return repo.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
		Mobile:   u.Mobile,
	})
}

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	u, err := repo.dao.FindByEmail(ctx, email)
	if err != nil {
		return &domain.User{}, err
	}
	return toDomainUser(u), nil
}
func toDomainUser(u dao.User) *domain.User {

	return &domain.User{
		Id:       u.Id,
		Mobile:   u.Mobile,
		Email:    u.Email,
		Password: u.Password,
	}
}
