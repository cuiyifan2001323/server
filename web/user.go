package web

import (
	"errors"
	"gin-study/common"
	"gin-study/domain"
	"gin-study/service"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
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

var (
	EmailConflictErr    = service.EmailConflictErr
	PasswordOrMobileErr = service.PasswordOrMobileErr
	//ErrUserNotFound  = service.ErrUserNotFound
)

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
	gr.POST("/edit", h.Edit)
	gr.POST("/profile", h.Profile)
}

func (h HandleUser) SignUp(ctx *gin.Context) {
	params := new(domain.User)
	err := ctx.ShouldBind(params)
	if err != nil {
		result := common.Err("参数错误", 200)
		ctx.JSON(200, result)
		return
	}
	isEmail, err := h.emailReg.MatchString(params.Email)
	isPhone, err := h.phoneReg.MatchString(strconv.Itoa(params.Mobile))
	if !isEmail || !isPhone || err != nil {
		result := common.Err("邮箱格式或者手机号格式不正确", 200)
		ctx.JSON(200, result)
		return
	}
	err = h.svc.Signup(ctx, domain.User{
		Email:    params.Email,
		Mobile:   params.Mobile,
		Password: params.Password,
	})
	var result common.Result
	switch {
	case err == nil:
		result = common.Success("注册成功", 200)
	case errors.Is(err, EmailConflictErr):
		result = common.Err(EmailConflictErr.Error(), 500)
	default:
		result = common.Err("系统错误", 500)
	}
	ctx.JSON(200, result)
}

func (h *HandleUser) Login(ctx *gin.Context) {
	params := new(domain.User)
	err := ctx.ShouldBind(params)
	if err != nil {
		result := common.Err("邮箱格式或者手机号格式不正确", 500)
		ctx.JSON(200, result)
		return
	}
	u, err := h.svc.Login(ctx, params)
	var result common.Result
	switch {
	case err == nil:
		sess := sessions.Default(ctx)
		sess.Set("id", u.Id)
		sess.Options(sessions.Options{
			// 15分钟
			MaxAge: int(time.Minute * 60 * 15),
		})
		err = sess.Save()
		if err != nil {
			result = common.Err("系统错误", 500)
			break

		}
		result = common.Success("登录成功", 200)
	case errors.Is(err, PasswordOrMobileErr):
		result = common.Err("密码或邮箱不正确", 500)
	default:
		result = common.Err("系统错误", 500)
	}
	ctx.JSON(200, result)
}
func (h *HandleUser) Edit(ctx *gin.Context) {
	params := new(domain.User)
	err := ctx.ShouldBind(params)
	if err != nil {
		result := common.Err("参数错误", 200)
		ctx.JSON(200, result)
		return
	}
	err = h.svc.Edit(ctx, params)
	var result common.Result
	switch {
	case err == nil:
		result = common.Success("修改用户信息成功", 200)
	default:
		result = common.Err("系统错误", 500)
	}
	//return
	ctx.JSON(200, result)
}

func (h *HandleUser) Profile(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		result := common.Err("参数错误", 500)
		ctx.JSON(200, result)
		return
	}
	userId, err := strconv.Atoi(id)
	if err != nil {
		result := common.Err("系统错误", 500)
		ctx.JSON(200, result)
		return
	}
	user, err := h.svc.GetUserInfo(ctx, int64(userId))
	result := common.Success(user, 200)
	ctx.JSON(200, result)
}
