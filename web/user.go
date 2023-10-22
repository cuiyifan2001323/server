package web

import (
	"gin-study/common"
	"gin-study/domain"
	"gin-study/service"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	emailReg = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	phoneReg = "^1[3456789][0-9]{9}$"
)

type HandleUser struct {
	emailReg *regexp.Regexp
	phoneReg *regexp.Regexp
	svc      *service.UserService
}

func NewUserHandle(svc *service.UserService) *HandleUser {
	return &HandleUser{
		emailReg: regexp.MustCompile(emailReg, regexp.None),
		phoneReg: regexp.MustCompile(phoneReg, regexp.None),
		svc:      svc,
	}
}
func (h *HandleUser) RegisterRoutes(server *gin.Engine) {
	gr := server.Group("/user")
	gr.POST("/signup", h.SignUp)
	gr.POST("/login", h.Login)
}

func (h HandleUser) SignUp(ctx *gin.Context) {
	params := new(domain.User)
	err := ctx.ShouldBind(params)
	isEmail, err := h.emailReg.MatchString(params.Email)
	isPhone, err := h.phoneReg.MatchString(strconv.Itoa(params.Mobile))
	if !isEmail || !isPhone || err != nil {
		result := common.Err("邮箱格式或者手机号格式不正确", 200)
		ctx.JSON(200, result)
		return
	}
	h.svc.Signup(ctx, domain.User{
		Email:    params.Email,
		Mobile:   params.Mobile,
		Password: params.Password,
	})
	result := common.Success("注册成功", 200)
	ctx.JSON(200, result)
}

func (h HandleUser) Login(server *gin.Context) {

}
