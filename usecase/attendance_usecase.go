package usecase

import (
	"TverClip/model"
	"TverClip/repository"
)

// IAttendanceUsecase インターフェースは勤怠データに関するビジネスロジックを定義します。
type IAttendanceUsecase interface {
	GetAttendanceByID(attendanceID uint) (model.Attendance, error)
	GetAttendanceByDate(date string, userID uint) (model.Attendance, error)
	FindOrCreateAttendanceByDate(date string, userID uint) (model.Attendance, error)
	CreateAttendance(attendance model.Attendance) (model.Attendance, error)
	UpdateAttendance(attendanceID uint, newAttendance model.Attendance) (model.Attendance, error)
	DeleteAttendance(attendanceID uint) error
}

// AttendanceUsecase 構造体はIAttendanceUsecaseインターフェースを実装します。
type AttendanceUsecase struct {
	ar repository.IAttendanceRepository
}

func NewAttendanceUsecase(ar repository.IAttendanceRepository) IAttendanceUsecase {
	return &AttendanceUsecase{ar}
}

func (au *AttendanceUsecase) GetAttendanceByDate(date string, userID uint) (model.Attendance, error) {
	attendance := model.Attendance{}
	if err := au.ar.GetAttendanceByDate(&attendance, date, userID); err != nil {
		return model.Attendance{}, err
	}
	return attendance, nil
}

func (au *AttendanceUsecase) GetAttendanceByID(attendanceID uint) (model.Attendance, error) {
	attendance := model.Attendance{}
	if err := au.ar.GetAttendanceByID(&attendance, attendanceID); err != nil {
		return model.Attendance{}, err
	}
	return attendance, nil
}

func (au *AttendanceUsecase) FindOrCreateAttendanceByDate(date string, userID uint) (model.Attendance, error) {
	attendance := model.Attendance{}
	if err := au.ar.FindOrCreateAttendanceByDate(&attendance, date, userID); err != nil {
		return model.Attendance{}, err
	}
	return attendance, nil
}

func (au *AttendanceUsecase) CreateAttendance(attendance model.Attendance) (model.Attendance, error) {
	createdAttendance, err := au.ar.CreateAttendance(&attendance)
	if err != nil {
		return model.Attendance{}, err
	}
	return createdAttendance, nil
}

func (au *AttendanceUsecase) UpdateAttendance(attendanceID uint, newAttendance model.Attendance) (model.Attendance, error) {
	updatedAttendance, err := au.ar.UpdateAttendance(attendanceID, &newAttendance)
	if err != nil {
		return model.Attendance{}, err
	}
	return updatedAttendance, nil
}

func (au *AttendanceUsecase) DeleteAttendance(attendanceID uint) error {
	if err := au.ar.DeleteAttendance(attendanceID); err != nil {
		return err
	}
	return nil
}
