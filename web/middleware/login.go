package middleware

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type LoginMiddlewareBuilder struct {
}

func (l *LoginMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.Request.URL.Path
		fmt.Println(path)
		if strings.HasPrefix(path, "/user/login") || strings.HasPrefix(path, "/user/signup") {
			return
		}
		sess := sessions.Default(context)
		if sess.Get("id") == nil {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
