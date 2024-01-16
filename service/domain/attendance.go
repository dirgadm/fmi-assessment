package domain

import (
	"context"
	"time"

	"github.com/dirgadm/fmi-assessment/service/domain/dto"
)

// Attendance is representing the absensi data struct
type Attendance struct {
	ID        int `gorm:"primaryKey;autoIncrement:true"`
	UserID    int
	Latitude  float64
	Longitude float64
	CreatedAt time.Time
}

func (m *Attendance) TableName() string {
	return "attendance"
}

// AttendanceUsecase represent the article's usecases
type AttendanceUsecase interface {
	CreateAttendance(ctx context.Context, req dto.AttendanceRequest) (res dto.AttendanceResponse, err error)
}

// AttendanceRepository represent the article's repository contract
type AttendanceRepository interface {
	Create(ctx context.Context, attendace *Attendance) (err error)
}
