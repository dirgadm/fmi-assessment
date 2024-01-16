package usecase

import (
	"context"
	"time"

	"github.com/dirgadm/fmi-assessment/pkg/constants"
	"github.com/dirgadm/fmi-assessment/service/domain"
	"github.com/dirgadm/fmi-assessment/service/domain/dto"
	"github.com/labstack/gommon/log"
)

type attendaceUsecase struct {
	attendaceRepo  domain.AttendanceRepository
	contextTimeout time.Duration
}

func NewAttendanceUsecase(u domain.AttendanceRepository, timeout time.Duration) domain.AttendanceUsecase {
	return &attendaceUsecase{
		attendaceRepo:  u,
		contextTimeout: timeout,
	}
}

func (u *attendaceUsecase) CreateAttendance(ctx context.Context, req dto.AttendanceRequest) (res dto.AttendanceResponse, err error) {
	userId := ctx.Value(constants.KeyUserID).(int)

	attendace := &domain.Attendance{
		UserID:    userId,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		CreatedAt: time.Now(),
	}
	err = u.attendaceRepo.Create(ctx, attendace)
	if err != nil {
		log.Error(err)
		return
	}

	res.Message = "Success Create Attendance"

	return
}
