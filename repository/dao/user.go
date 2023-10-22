package dao

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}
type User struct {
	Id int64 `gorm:"primaryKey"`
	// 邮箱
	Email string
	// 密码
	Password string
	// 手机号
	Mobile    int
	UpdatedAt int64 `gorm:"column:update_time"`
	CreatedAt int64 `gorm:"column:create_time"`
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	dsn := "root:cyf2001323@tcp(127.0.0.1:13316)/study?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db.WithContext(ctx).Create(&u).Error
}
