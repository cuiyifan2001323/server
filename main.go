package main

import (
	"fmt"
	"gin-study/repository"
	"gin-study/repository/dao"
	"gin-study/service"
	"gin-study/web"
	"gin-study/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

func main() {
	g := gin.Default()
	g.Use(cors.New(cors.Config{
		// 业务请求中可以带上的头
		AllowHeaders: []string{"Content-Type"},
		// 是否允许带上用户认证 信息（比如 cookie）。
		AllowCredentials: true,
		// 哪些来源是允许的
		AllowOriginFunc: func(origin string) bool {
			return strings.HasPrefix(origin, "http://localhost")
		},
	}))
	store := cookie.NewStore([]byte("secret"))
	g.Use(sessions.Sessions("ssid", store))
	loginMiddleware := &middleware.LoginMiddlewareBuilder{}
	g.Use(loginMiddleware.CheckLogin())
	db := InitDb()
	InitUserHandle(db, g)
	g.Run(":8080")
}
func InitDb() *gorm.DB {
	dsn := "root:cyf2001323@tcp(127.0.0.1:13316)/study?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("数据库连接失败")
	}
	return db
}
func InitUserHandle(db *gorm.DB, g *gin.Engine) {
	userDao := dao.NewUserDao(db)
	userRepo := repository.NewUserRepository(userDao)
	sve := service.NewUserService(userRepo)
	userRouter := web.NewUserHandle(sve)
	userRouter.RegisterRoutes(g)
}
