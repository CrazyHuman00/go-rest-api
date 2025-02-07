package router

import (
	"go-rest-api/controller"

	"github.com/labstack/echo/v4"
)

// コンストラクタ（controllerとのDI）
func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)  // singupのエンドポイントを追加
	e.POST("/login", uc.LogIn)    // loginのエンドポイントを追加
	e.POST("/logout", uc.LogOut)  // logoutのエンドポイントを追加
	return e
}