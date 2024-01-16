package mysql

import (
	"context"

	"github.com/dirgadm/fmi-assessment/service/domain"
	"gorm.io/gorm"
)

type uploadRepository struct {
	Conn *gorm.DB
}

func NewlUploadRepository(conn *gorm.DB) domain.UploadRepository {
	return &uploadRepository{conn}
}

func (m *uploadRepository) GetByUserId(ctx context.Context, offset int, limit int, search string, userId int) (upload []domain.UploadedFile, count int64, err error) {
	gorm := m.Conn.Model(domain.UploadedFile{})

	if userId != 0 {
		gorm = gorm.Where("user_id = ?", userId)
	}

	err = gorm.Count(&count).Error
	if err != nil {
		return
	}

	err = gorm.Offset(offset).Limit(limit).Find(&upload).Error

	return
}

func (m *uploadRepository) DeleteByUserId(ctx context.Context, listId int) (err error) {
	err = m.Conn.Where("user_id = ?", listId).Delete(domain.UploadedFile{}).Error
	return
}

func (m *uploadRepository) Create(ctx context.Context, upload *domain.UploadedFile) (err error) {
	err = m.Conn.Create(upload).Error
	return
}
