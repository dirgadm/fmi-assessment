package mysql

import (
	"context"

	"github.com/dirgadm/fmi-assessment/service/domain"
	"gorm.io/gorm"
)

type attendaceRepository struct {
	Conn *gorm.DB
}

func NewlAttendanceRepository(conn *gorm.DB) domain.AttendanceRepository {
	return &attendaceRepository{conn}
}

// Create implements domain.AbsensiRepository.
func (m *attendaceRepository) Create(ctx context.Context, absensi *domain.Attendance) (err error) {
	result := m.Conn.Exec("CALL absensi(?, ?, ?)", absensi.UserID, absensi.Latitude, absensi.Longitude)

	return result.Error
}
