package service

import (
	"context"
	"errors"
	"gin-study/domain"
	"gin-study/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	repo *repository.UserRepository
}

var (
	// 密码错误
	PasswordOrMobileErr = errors.New("密码或手机号不正确")
	// 邮箱已存在错误
	EmailConflictErr = repository.EmailConflictErr
	// 没查到这个用户( 但是为了安全性，不会提示的这么明显 )
	//ErrUserNotFound = repository.ErrUserNotFound
)

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
func (s UserService) Signup(ctx context.Context, u domain.User) error {
	user, err := s.repo.FindByEmail(ctx, u.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	// 如果查询到了表示邮箱已存在了，不能在注册
	if user.Email != "" {
		return EmailConflictErr
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return s.repo.Create(ctx, u)
}

func (s UserService) Login(ctx context.Context, user *domain.User) (*domain.User, error) {
	u, err := s.repo.FindByEmail(ctx, user.Email)
	if err != nil {
		// 如果是没找到的error，那么返回PasswordOrMobileErr，提示账号或密码不正确
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return u, PasswordOrMobileErr
		}
		return u, err
	}
	isPassword := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if isPassword != nil {
		return u, PasswordOrMobileErr
	}
	return u, nil
}

func (s UserService) Edit(ctx *gin.Context, userInfo *domain.User) error {
	_, err := s.repo.FindById(ctx, userInfo.Id)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, userInfo)
}

func (s UserService) GetUserInfo(ctx *gin.Context, id int64) (*domain.User, error) {
	return s.repo.FindById(ctx, id)
}
