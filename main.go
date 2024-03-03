package main

import (
	"TverClip/controller"
	"TverClip/db"
	"TverClip/repository"
	"TverClip/router"
	"TverClip/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	attendanceRepository := repository.NewAttendanceRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	attendanceUsecase := usecase.NewAttendanceUsecase(attendanceRepository)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	ateendanceController := controller.NewAttendanceController(attendanceUsecase)
	e := router.NewRouter(userController, taskController, ateendanceController)
	e.Logger.Fatal(e.Start("0.0.0.0:8083"))
}
