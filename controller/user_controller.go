package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"

	"go-rest-api/model"
	"go-rest-api/usecase"
)

// controllerのインタフェースを定義
type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
}

// controllerの構造体を定義
type userController struct {
	uu usecase.IUserUsecase
}

// コンストラクタ（usecaseとのDI）
func NewUserController(uu usecase.IUserUsecase) IUserController {
	// 構造体の実体のポインタを返す
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	// クライアントから送られたリクエストボディを構造体に変換
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// バインドに成功した場合usecaseのSignUpメソッドを呼び出す
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) LogIn(c echo.Context) error {
	// クライアントから送られたリクエストボディを構造体に変換
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// バインドに成功した場合usecaseのLoginメソッドを呼び出す
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// サーバサイドにクッキーで設定
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	// cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode // クロスドメイン間でのCookieの送信を許可
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) LogOut(c echo.Context) error {
	// クッキーの削除
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = "" // 空文字にすることで削除される
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	// cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode // クロスドメイン間でのCookieの送信を許可
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}