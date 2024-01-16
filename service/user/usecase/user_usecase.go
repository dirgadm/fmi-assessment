package usecase

import (
	"context"
	"time"

	"github.com/dirgadm/fmi-assessment/pkg/ehttp"
	jwt "github.com/dirgadm/fmi-assessment/pkg/jwt"
	"github.com/dirgadm/fmi-assessment/service/domain"
	"github.com/dirgadm/fmi-assessment/service/domain/dto"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo domain.UserRepository
	// authorRepo     domain.AuthorRepository
	contextTimeout time.Duration
}

// NewUserUsecase will create new an articleUsecase object representation of domain.UserUsecase interface
func NewUserUsecase(u domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       u,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) Login(ctx context.Context, req dto.LoginRequest) (res dto.AuthResponse, err error) {
	var user domain.User
	user, err = u.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		log.Error(err)
		err = ehttp.ErrorOutput("email", "The email is not match")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Error(err)
		err = ehttp.ErrorOutput("password", "The password is not match")
		return
	}

	jwtInit := jwt.NewJWT([]byte(viper.GetString(`jwt.key`)))
	expiredAt := time.Now().Add(time.Hour * 1)
	claim := jwt.UserClaim{
		UserId:    user.Id,
		ExpiresAt: expiredAt.Unix(),
	}

	token, err := jwtInit.Create(claim)
	if err != nil {
		log.Error(err)
		return
	}

	us := &domain.User{
		Id:        user.Id,
		Email:     user.Email,
		Password:  user.Password,
		Name:      user.Name,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: time.Now(),
	}
	if err = u.userRepo.Update(ctx, us); err != nil {
		log.Error(err)
		return
	}

	res.Token = token
	res.ExpiredAt = expiredAt
	res.User = dto.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return
}

func (u *userUsecase) Register(ctx context.Context, req dto.RegisterRequest) (res dto.AuthResponse, err error) {
	var userExist domain.User
	userExist, _ = u.userRepo.GetByEmail(ctx, req.Email)
	if userExist.Id != 0 {
		log.Error(err)
		err = ehttp.ErrorOutput("email", "The email is already exists")
		return
	}
	if req.Phone[:3] != "+62" {
		log.Error(err)
		err = ehttp.ErrorOutput("phone", "The phone number should be start with +62")
		return
	}

	var bytes []byte
	bytes, err = bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		log.Error(err)
		return
	}
	passwordHash := string(bytes)

	user := &domain.User{
		Email:     req.Email,
		Password:  passwordHash,
		Name:      req.Name,
		Phone:     req.Phone,
		UpdatedAt: time.Now(),
	}

	if req.Photo == "" {
		err = u.userRepo.CreateWithoutPhoto(ctx, user)
		if err != nil {
			log.Error(err)
			return
		}
	} else {
		err = u.userRepo.CreateWithPhoto(ctx, user, req.Photo)
		if err != nil {
			log.Error(err)
			return
		}
	}

	jwtInit := jwt.NewJWT([]byte(viper.GetString(`jwt.key`)))
	expiredAt := time.Now().Add(time.Hour * 1)
	claim := jwt.UserClaim{
		UserId:    user.Id,
		ExpiresAt: expiredAt.Unix(),
	}

	token, err := jwtInit.Create(claim)
	if err != nil {
		log.Error(err)
		return
	}

	res.Token = token
	res.ExpiredAt = expiredAt
	res.User = dto.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return
}
