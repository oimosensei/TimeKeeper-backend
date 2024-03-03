package repository

import (
	"TverClip/model"
	"fmt"

	"gorm.io/gorm"
)

type IAttendanceRepository interface {
	GetAttendanceByID(attendance *model.Attendance, attendanceID uint) error
	GetAttendanceByDate(attendance *model.Attendance, date string, userID uint) error
	FindOrCreateAttendanceByDate(attendance *model.Attendance, date string, userID uint) error
	CreateAttendance(attendance *model.Attendance) (model.Attendance, error)
	UpdateAttendance(attendanceID uint, newAttendance *model.Attendance) (model.Attendance, error)
	DeleteAttendance(attendanceID uint) error
}

type AttendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) IAttendanceRepository {
	return &AttendanceRepository{db}
}

func (ar *AttendanceRepository) GetAttendanceByDate(attendance *model.Attendance, date string, userID uint) error {
	if err := ar.db.Where("date = ? AND user_id = ?", date, userID).First(attendance).Error; err != nil {
		return err
	}
	return nil
}

func (ar *AttendanceRepository) GetAttendanceByID(attendance *model.Attendance, attendanceID uint) error {
	if err := ar.db.Where("id = ?", attendanceID).First(attendance).Error; err != nil {
		return err
	}
	return nil
}

func (ar *AttendanceRepository) FindOrCreateAttendanceByDate(attendance *model.Attendance, date string, userID uint) error {
	// 指定された日付とユーザーIDで出席データを検索
	err := ar.db.Where("date = ? AND user_id = ?", date, userID).First(attendance).Error
	if err == gorm.ErrRecordNotFound {
		// レコードが見つからない場合、新しい勤怠データを作成
		attendance.Date = date
		attendance.UserID = userID

		err = ar.db.Create(attendance).Error
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

func (ar *AttendanceRepository) CreateAttendance(attendance *model.Attendance) (model.Attendance, error) {
	if err := ar.db.Create(attendance).Error; err != nil {
		return model.Attendance{}, err
	}
	return *attendance, nil
}

func (ar *AttendanceRepository) UpdateAttendance(attendanceID uint, newAttendance *model.Attendance) (model.Attendance, error) {
	result := ar.db.Model(&model.Attendance{}).Where("id = ?", attendanceID).Updates(newAttendance)

	if result.Error != nil {
		return model.Attendance{}, result.Error
	}

	if result.RowsAffected < 1 {
		return model.Attendance{}, fmt.Errorf("object does not exist")
	}

	return *newAttendance, nil
}

func (ar *AttendanceRepository) DeleteAttendance(attendanceID uint) error {
	result := ar.db.Where("id = ?", attendanceID).Delete(&model.Attendance{})

	if result.Error != nil {
		return result.Error
	}
	return nil
}
