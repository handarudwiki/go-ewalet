package service

import (
	"bytes"
	"context"
	"errors"
	"text/template"
	"time"

	"github.com/handarudwiki/golang-ewalet/domain"
	"github.com/handarudwiki/golang-ewalet/dto"
)

type NotificationService struct {
	notificationRepository domain.NotificationRepository
	templateRepository     domain.TemplateRepository
	hub                    *dto.Hub
}

func NewNotification(notificationRepository domain.NotificationRepository, templateRepository domain.TemplateRepository, hub *dto.Hub) domain.NotificationService {
	return &NotificationService{
		notificationRepository: notificationRepository,
		templateRepository:     templateRepository,
		hub:                    hub,
	}
}

func (s NotificationService) FindByUser(ctx context.Context, user int64) ([]dto.NotificationData, error) {
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

func (s NotificationService) Insert(ctx context.Context, userId int64, code string, data map[string]string) error {
	tmpl, err := s.templateRepository.FindByCode(ctx, code)

	if err != nil {
		return err
	}

	if tmpl == (domain.Template{}) {
		return errors.New("error template not found")
	}

	body := new(bytes.Buffer)
	t := template.Must(template.New("notif").Parse(tmpl.Body))
	err = t.Execute(body, data)

	if err != nil {
		return err
	}
	notification := domain.Notification{
		UserID:    userId,
		Title:     tmpl.Title,
		Body:      body.String(),
		Status:    1,
		IsRead:    0,
		CreatedAt: time.Now(),
	}

	err = s.notificationRepository.Insert(ctx, &notification)
	if err != nil {
		return err
	}
	if channel, ok := s.hub.NotificationChannel[userId]; ok {
		channel <- dto.NotificationData{
			ID:        notification.ID,
			Title:     notification.Title,
			Body:      notification.Body,
			Status:    notification.Status,
			IsRead:    notification.IsRead,
			CreatedAt: notification.CreatedAt,
		}
	}

	return nil
}
