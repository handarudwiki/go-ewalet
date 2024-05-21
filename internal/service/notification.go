package service

import (
	"context"

	"github.com/handarudwiki/golang-ewalet/domain"
	"github.com/handarudwiki/golang-ewalet/dto"
)

type NotificationService struct {
	notificationRepository domain.NotificationRepository
}

func NewNotification(notificationRepository domain.NotificationRepository) domain.NotificationService {
	return &NotificationService{
		notificationRepository: notificationRepository,
	}
}

func (s *NotificationService) FindByUser(ctx context.Context, user int64) ([]dto.NotificationData, error) {
	notifications, err := s.notificationRepository.FindByUser(ctx, user)

	if err != nil {
		return nil, err
	}

	var results []dto.NotificationData

	for _, notification := range notifications {
		results = append(results, dto.NotificationData{
			ID:        notification.ID,
			Title:     notification.Title,
			Body:      notification.Title,
			Status:    notification.Status,
			IsRead:    notification.IsRead,
			CreatedAt: notification.CreatedAt,
		})
	}

	if results == nil {
		results = make([]dto.NotificationData, 0)
	}

	return results, nil
}
