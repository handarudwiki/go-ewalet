package domain

import (
	"context"
	"time"

	"github.com/handarudwiki/golang-ewalet/dto"
)

type Notification struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	Status    int8      `db:"status"`
	IsRead    int8      `db:"is_read"`
	CreatedAt time.Time `db:"created_at"`
}

type NotificationRepository interface {
	FindByUser(ctx context.Context, user int64) ([]Notification, error)
	Insert(ctx context.Context, notification *Notification) error
	Update(ctx context.Context, notification *Notification) error
}

type NotificationService interface {
	FindByUser(ctx context.Context, user int64) ([]dto.NotificationData, error)
}
