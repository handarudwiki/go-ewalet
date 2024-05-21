package repository

import (
	"context"

	"github.com/handarudwiki/golang-ewalet/domain"
	"gorm.io/gorm"
)

type notificationRepository struct {
	db *gorm.DB
}

func NewNotification(db *gorm.DB) domain.NotificationRepository {
	return &notificationRepository{
		db: db,
	}
}

func (r *notificationRepository) FindByUser(ctx context.Context, user int64) ([]domain.Notification, error) {
	var notifications []domain.Notification

	err := r.db.Where("user_id=?", user).Find(&notifications).Error

	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func (r *notificationRepository) Insert(ctx context.Context, notification *domain.Notification) error {
	err := r.db.Create(&notification).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *notificationRepository) Update(ctx context.Context, notification *domain.Notification) error {
	var data domain.Notification

	err := r.db.Where("id=?", notification.ID).Error

	if err != nil {
		return err
	}

	data.Title = notification.Title
	data.Body = notification.Body
	data.Status = notification.Status
	data.IsRead = notification.IsRead

	err = r.db.Save(&data).Error

	if err != nil {
		return err
	}

	return nil
}
