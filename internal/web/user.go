package web

import (
	"errors"
	"github.com/gin-gonic/gin"
	"walnut/exception"
	"walnut/internal/service"
	"walnut/response"
)

type UserHandler struct {
	srv service.UserSrv
}

func NewUserHandler(srv service.UserSrv) *UserHandler {
	return &UserHandler{srv: srv}
}

func (uh *UserHandler) RegisterRoutes(router gin.IRoutes) {

	router.POST("/logout", uh.Logout)
	router.POST("/login", uh.Login)
	router.POST("/register", uh.Register)
}

func (uh *UserHandler) Login(ctx *gin.Context) {
	uh.srv.Login()
	data := map[string]string{
		"token": "123",
		"uid":   "123",
	}
	response.Success(ctx, data)
}
func (uh *UserHandler) Logout(ctx *gin.Context) {
	uh.srv.Logout()

	response.Failed(ctx, exception.ErrValidateFailed("请求体结构错误,err:%s", "username 长度大于10"))
}
func (uh *UserHandler) Register(ctx *gin.Context) {
	uh.srv.Register()

	response.Failed(ctx, errors.New("注册失败"))
}
