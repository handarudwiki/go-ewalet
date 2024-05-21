package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/handarudwiki/golang-ewalet/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUser(cons *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: cons,
	}

}
func (u userRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	var user domain.User
	err := u.db.Where("id=?", id).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
func (u userRepository) FindByUsername(ctx context.Context, username string) (domain.User, error) {
	var user domain.User
	err := u.db.Where("username=?", username).First(&user).Error
	if err != nil {
		return domain.User{}, domain.ErrUserNotFound
	}
	return user, nil
}

func (u userRepository) Insert(ctx context.Context, user *domain.User) error {
	err := u.db.Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (u userRepository) Update(ctx context.Context, user *domain.User, verifiedAt time.Time) error {
	user.EmailVerifiedAt = sql.NullTime{
		Time:  verifiedAt,
		Valid: true,
	}
	var data domain.User

	err := u.db.Where("id=?", user.ID).First(&data).Error

	if err != nil {
		return err
	}

	data.Email = user.Email
	data.Name = user.Name
	data.EmailVerifiedAt = user.EmailVerifiedAt
	data.Phone = user.Phone
	data.Password = user.Password
	data.Username = user.Username

	err = u.db.Save(&data).Error

	if err != nil {
		return err
	}

	return nil
}
