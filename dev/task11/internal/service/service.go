package service

import (
	"context"
	"github.com/infamax/l2/task11/internal/models"
	"github.com/infamax/l2/task11/internal/repository"
)

type Service interface {
	CreateEvent(ctx context.Context, event *models.Event) (int, error)
	UpdateEvent(ctx context.Context, event *models.Event) error
	DeleteEvent(ctx context.Context, eventID int) error
	GetEventsForDay(ctx context.Context, userID int, date string) ([]models.Event, error)
	GetEventsForWeek(ctx context.Context, userID int, date string) ([]models.Event, error)
	GetEventsForMonth(ctx context.Context, userID int, date string) ([]models.Event, error)
}

type service struct {
	repo repository.Repository
}

func New(repo repository.Repository) (*service, error) {
	if repo == nil {
		return nil, nil
	}

	return &service{
		repo: repo,
	}, nil
}
