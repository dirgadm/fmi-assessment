package domain

import (
	"context"
	"net/http"
	"time"

	"github.com/dirgadm/fmi-assessment/service/domain/dto"
)

type UploadedFile struct {
	Id        int `gorm:"primaryKey;autoIncrement:true"`
	FileName  string
	UserId    int
	CreatedAt time.Time
}

func (m *UploadedFile) TableName() string {
	return "uploaded_files"
}

type UploadUsecase interface {
	UploadFile(ctx context.Context, w http.ResponseWriter, r *http.Request) (res []dto.UploadResponse, err error)
}

type UploadRepository interface {
	GetByUserId(ctx context.Context, offset int, limit int, search string, userId int) (uf []UploadedFile, count int64, err error)
	DeleteByUserId(ctx context.Context, userId int) (err error)
	Create(ctx context.Context, upload *UploadedFile) (err error)
}
