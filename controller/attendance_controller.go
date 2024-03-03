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

// IAttendanceController インターフェースは勤怠管理のためのコントローラーが実装すべきメソッドを定義します。
type IAttendanceController interface {
	GetAttendanceByID(c echo.Context) error
	GetAttendanceByDate(c echo.Context) error
	FindOrCreateAttendanceByDate(c echo.Context) error
	CreateAttendance(c echo.Context) error
	UpdateAttendance(c echo.Context) error
	DeleteAttendance(c echo.Context) error
}

// AttendanceController 構造体はIAttendanceControllerインターフェースを実装し、
// 勤怠データに関する操作を行います。
type AttendanceController struct {
	au usecase.IAttendanceUsecase
}

// NewAttendanceController はAttendanceControllerの新しいインスタンスを作成します。
func NewAttendanceController(au usecase.IAttendanceUsecase) *AttendanceController {
	return &AttendanceController{au}
}

func (ac *AttendanceController) GetAttendanceByID(c echo.Context) error {
	attendanceIdString := c.Param("attendanceId")
	attendanceId, _ := strconv.Atoi(attendanceIdString)
	attendance, err := ac.au.GetAttendanceByID(uint(attendanceId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, attendance)
}

// getAttendasnceByDate
func (ac *AttendanceController) GetAttendanceByDate(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	userId := claims["user_id"]

	date := c.Param("date")
	attendance, err := ac.au.GetAttendanceByDate(date, uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, attendance)
}

func (ac *AttendanceController) FindOrCreateAttendanceByDate(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	userId := claims["user_id"]

	date := c.Param("date")
	attendance, err := ac.au.FindOrCreateAttendanceByDate(date, uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, attendance)
}

func (ac *AttendanceController) CreateAttendance(c echo.Context) error {
	attendance := model.Attendance{}
	if err := c.Bind(&attendance); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	attendanceRes, err := ac.au.CreateAttendance(attendance)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, attendanceRes)
}

func (ac *AttendanceController) UpdateAttendance(c echo.Context) error {
	attendanceIdString := c.Param("attendanceId")
	attendanceId, _ := strconv.Atoi(attendanceIdString)
	attendance := model.Attendance{}
	if err := c.Bind(&attendance); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	updatedAttendance, err := ac.au.UpdateAttendance(uint(attendanceId), attendance)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedAttendance)
}

func (ac *AttendanceController) DeleteAttendance(c echo.Context) error {
	attendanceIdString := c.Param("attendanceId")
	attendanceId, _ := strconv.Atoi(attendanceIdString)
	if err := ac.au.DeleteAttendance(uint(attendanceId)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
