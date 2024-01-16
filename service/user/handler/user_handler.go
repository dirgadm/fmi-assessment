package http

import (
	"github.com/dirgadm/fmi-assessment/pkg/ehttp"
	"github.com/dirgadm/fmi-assessment/service/domain"
	"github.com/dirgadm/fmi-assessment/service/domain/dto"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// UserHandler  represent the httphandler for User
type UserHandler struct {
	PUsecase domain.UserUsecase
}

// NewUserHandler will initialize the User resources endpoint
func NewUserHandler(e *echo.Echo, ps domain.UserUsecase) {
	handler := &UserHandler{
		PUsecase: ps,
	}
	v1 := e.Group("v1")

	v1.POST("/auth/login", handler.Login)
	v1.POST("/auth/register", handler.Register)
}

func (h UserHandler) Login(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)
	validator := validator.New()

	var req dto.LoginRequest

	if err = ctx.Bind(&req); err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	if err = validator.Struct(req); err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	ctx.ResponseData, err = h.PUsecase.Login(ctx.Request().Context(), req)

	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	return ctx.Serve(err)
}

func (h UserHandler) Register(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)
	validator := validator.New()

	var req dto.RegisterRequest

	if err = ctx.Bind(&req); err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	if err = validator.Struct(req); err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	ctx.ResponseData, err = h.PUsecase.Register(ctx.Request().Context(), req)
	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	return ctx.Serve(err)
}
