package router

import (
	"go-rest-api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// コンストラクタ（controllerとのDI）
func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()
	// ミドルウェアの設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// フロントエンドのドメインを指定、vercelにデプロイしたときに取得するドメインをFE_URLに設定
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{
			echo.HeaderOrigin, 
			echo.HeaderContentType, 
			echo.HeaderAccept, 
			echo.HeaderAccessControlAllowHeaders, 
			echo.HeaderXCSRFToken},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true, // クッキーを許可
	}))

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath: "/",
		CookieDomain: os.Getenv("FE_URL"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode, 
		// CookieSameSite: http.SameSiteDefaultMode,
		// CookieMaxAge: 60, // 1分間有効
	}))

	// ユーザのエンドポイントを追加
	e.POST("/signup", uc.SignUp)  // singupのエンドポイントを追加
	e.POST("/login", uc.LogIn)    // loginのエンドポイントを追加
	e.POST("/logout", uc.LogOut)  // logoutのエンドポイントを追加
	e.GET("/csrf", uc.CsrfToken)  // csrfのエンドポイントを追加

	t := e.Group("/tasks")        // tasksのグループを作成
	// tasksグループに対してJWTのミドルウェアを適用
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey : []byte(os.Getenv("SECRET")), // jwtの署名キーを環境変数から取得
		TokenLookup: "cookie:token",              // ログインユーザの情報を取得するためのキーを指定
	}))
	t.GET("", tc.GetAllTasks)           // tasksのGetAllTasksのエンドポイントを追加
	t.GET("/:taskId", tc.GetTaskById)   // tasksのGetTaskByIdのエンドポイントを追加
	t.POST("", tc.CreateTask)           // tasksのCreateTaskのエンドポイントを追加
	t.PUT("/:taskId", tc.UpdateTask)    // tasksのUpdateTaskのエンドポイントを追加
	t.DELETE("/:taskId", tc.DeleteTask) // tasksのDeleteTaskのエンドポイントを追加
	return e
}