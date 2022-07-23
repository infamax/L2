package service

import (
	"context"
	"github.com/infamax/l2/task11/internal/models"
)

func (s *service) CreateEvent(ctx context.Context, event *models.Event) (int, error) {
	return s.repo.CreateEvent(ctx, event)
}
