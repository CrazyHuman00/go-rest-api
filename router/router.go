package router

import (
	"go-rest-api/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// コンストラクタ（controllerとのDI）
func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)  // singupのエンドポイントを追加
	e.POST("/login", uc.LogIn)    // loginのエンドポイントを追加
	e.POST("/logout", uc.LogOut)  // logoutのエンドポイントを追加
	t := e.Group("/tasks")        // tasksのグループを作成
	// tasksグループに対してJWTのミドルウェアを適用
	t.Use(echojwt.WithConfig(echojwt.Config{
		// jwtの署名キーを環境変数から取得
		SigningKey : []byte(os.Getenv("SECRET")),
		// ログインユーザの情報を取得するためのキーを指定
		TokenLookup: "cookie:token",
	}))

	t.GET("", tc.GetAllTasks)           // tasksのGetAllTasksのエンドポイントを追加
	t.GET("/:taskId", tc.GetTaskById)   // tasksのGetTaskByIdのエンドポイントを追加
	t.POST("", tc.CreateTask)           // tasksのCreateTaskのエンドポイントを追加
	t.PUT("/:taskId", tc.UpdateTask)    // tasksのUpdateTaskのエンドポイントを追加
	t.DELETE("/:taskId", tc.DeleteTask) // tasksのDeleteTaskのエンドポイントを追加
	return e
}