package dao

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

// 预定义一个错误( 邮箱重复错误 )
var EmailConflictErr = errors.New("此邮箱已存在，请换一个")

type User struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	// 邮箱
	Email string `gorm:"unique"`
	// 密码
	Password string
	// 手机号
	Mobile    int   `gorm:"unique"`
	UpdatedAt int64 `gorm:"column:update_time"`
	CreatedAt int64 `gorm:"column:create_time"`
}

func NewUserDao(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}
func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	err := dao.db.WithContext(ctx).Create(&u).Error
	return err
	// 断言是不是mysql错误
	//if errNumber, ok := err.(*mysql.MySQLError); ok {
	//	// mysql 错误状态吗 1062是主键重复了，注册接口里面只有邮箱是主键
	//	if errNumber.Number == 1062 {
	//		return EmailConflictErr
	//	}
	//}
}

func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := dao.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return user, err
}
