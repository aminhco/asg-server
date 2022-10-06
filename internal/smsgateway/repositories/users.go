package repositories

import (
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/models"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound = gorm.ErrRecordNotFound
)

type UsersRepository struct {
	db *gorm.DB
}

func (r *UsersRepository) GetByLogin(login string) (models.User, error) {
	user := models.User{}

	return user, r.db.Where("id = ?", login).Preload("Devices").Take(&user).Error
}

func (r *UsersRepository) Insert(user *models.User) error {
	return r.db.Create(user).Error
}

func NewUsersRepository(db *gorm.DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}
