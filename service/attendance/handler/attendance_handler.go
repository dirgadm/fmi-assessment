package http

import (
	"github.com/dirgadm/fmi-assessment/pkg/ehttp"
	"github.com/dirgadm/fmi-assessment/pkg/middleware"
	"github.com/dirgadm/fmi-assessment/service/domain"
	"github.com/dirgadm/fmi-assessment/service/domain/dto"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// AttendanceHandler  represent the httphandler for Attendance
type AttendanceHandler struct {
	PUsecase domain.AttendanceUsecase
}

// NewAttendanceHandler will initialize the Attendance resources endpoint
func NewAttendanceHandler(e *echo.Echo, ps domain.AttendanceUsecase) {
	handler := &AttendanceHandler{
		PUsecase: ps,
	}
	v1 := e.Group("v1")

	mw := middleware.NewMiddleware()

	v1.POST("/attendance", handler.Create, mw.Authorized())
}

func (h AttendanceHandler) Create(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)
	validator := validator.New()

	var req dto.AttendanceRequest

	if err = ctx.Bind(&req); err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	if err = validator.Struct(req); err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	ctx.ResponseData, err = h.PUsecase.CreateAttendance(ctx.Request().Context(), req)

	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	return ctx.Serve(err)
}
