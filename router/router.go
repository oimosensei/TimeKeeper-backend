package router

import (
	"TverClip/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController, ac controller.IAttendanceController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL"), "http://127.0.0.1:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken, echo.HeaderAccessControlAllowOrigin},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/clips", tc.GetAllTasks)
	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))

	t.GET("/:taskId", tc.GetTaskByID)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)

	// Attendanceに関するルートのグループ
	a := e.Group("/attendances")
	a.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	// Attendanceデータに関するエンドポイント
	a.GET("/:attendanceId", ac.GetAttendanceByID)          // 特定のAttendanceをIDで取得
	a.POST("", ac.CreateAttendance)                        // 新しいAttendanceを作成
	a.PUT("/:attendanceId", ac.UpdateAttendance)           // 特定のAttendanceを更新
	a.DELETE("/:attendanceId", ac.DeleteAttendance)        // 特定のAttendanceを削除
	a.GET("/date/:date", ac.GetAttendanceByDate)           // 特定の日付のAttendanceを取得
	a.POST("/date/:date", ac.FindOrCreateAttendanceByDate) // 特定の日付のAttendanceを取得

	s := e.Group("/session")
	s.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	s.GET("/user", uc.GetUserBySession)
	return e
}
