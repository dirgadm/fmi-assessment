package mysql

import (
	"context"

	"github.com/dirgadm/fmi-assessment/service/domain"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) domain.UserRepository {
	return &mysqlUserRepository{conn}
}

func (m *mysqlUserRepository) GetDetail(ctx context.Context, id int) (user domain.User, err error) {
	err = m.Conn.Where("id = ?", id).First(&user).Error
	return
}

func (m *mysqlUserRepository) GetByEmail(ctx context.Context, email string) (user domain.User, err error) {
	err = m.Conn.Where("email = ?", email).First(&user).Error
	return
}

func (m *mysqlUserRepository) CreateWithPhoto(ctx context.Context, user *domain.User, userPhoto string) (err error) {
	return m.Conn.Exec("CALL register_user_with_photo(?, ?, ?, ?, ?)", user.Email, user.Password, user.Name, user.Phone, userPhoto).Error
}

func (m *mysqlUserRepository) CreateWithoutPhoto(ctx context.Context, user *domain.User) (err error) {
	return m.Conn.Exec("CALL register_user(?, ?, ?, ?)", user.Email, user.Password, user.Name, user.Phone).Error
}

func (m *mysqlUserRepository) Update(ctx context.Context, user *domain.User) (err error) {
	err = m.Conn.Save(user).Error
	return
}
