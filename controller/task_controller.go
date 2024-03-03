package controller

import (
	"TverClip/model"
	"TverClip/usecase"
	"fmt"
	"net/http"
	"strconv"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTasksByUserID(c echo.Context) error
	GetTaskByID(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type TaskController struct {
	tu usecase.ITaskUsecase
}

func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &TaskController{tu}
}

// getalltasks
func (tc *TaskController) GetAllTasks(c echo.Context) error {
	println("getalltasks")
	tasks, err := tc.tu.GetAllTasks()
	if err != nil {
		println(err.Error())
		println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTasksByUserID(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	userId := claims["user_id"]

	tasks, err := tc.tu.GetTasksByUserID(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTaskByID(c echo.Context) error {
	taskIdString := c.Param("taskId")

	taskId, _ := strconv.Atoi(taskIdString)
	task, err := tc.tu.GetTaskByID(uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, task)
}

func (tc *TaskController) CreateTask(c echo.Context) error {
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		fmt.Print(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Print(task)

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	userId := uint(claims["user_id"].(float64))
	task.UserID = userId

	taskRes, err := tc.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, taskRes)
}

func (tc *TaskController) UpdateTask(c echo.Context) error {
	taskIDString := c.Param("taskId")
	taskID, _ := strconv.Atoi(taskIDString)
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	userId := uint(claims["user_id"].(float64))

	taskRes, err := tc.tu.UpdateTask(uint(taskID), userId, task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, taskRes)
}

func (tc *TaskController) DeleteTask(c echo.Context) error {
	taskIDString := c.Param("taskId")
	taskId, _ := strconv.Atoi(taskIDString)

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	userId := uint(claims["user_id"].(float64))

	if err := tc.tu.DeleteTask(uint(taskId), userId); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
