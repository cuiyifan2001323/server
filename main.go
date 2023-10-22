package main

import (
	"gin-study/service"
	"gin-study/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
)

//var DB *gorm.DB

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
	//dsn := "root:cyf2001323@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//DB = db
	//if err != nil {
	//	fmt.Println(err)
	//}
	sve := new(service.UserService)
	userRouter := web.NewUserHandle(sve)
	userRouter.RegisterRoutes(g)
	g.Run(":8080")

	//// 路径参数
	//g.GET("/users/:id", func(ctx *gin.Context) {
	//	id := ctx.Param("id")
	//	ctx.JSON(http.StatusOK, id)
	//})
	//// 通配符
	//g.GET("/file/*.html", func(context *gin.Context) {
	//	view := context.Param(".html")
	//	context.JSON(200, view)
	//})
}
