package service

import (
	"context"
	"github.com/infamax/l2/task11/internal/models"
)

func (s *service) UpdateEvent(ctx context.Context, event *models.Event) error {
	return s.repo.UpdateEvent(ctx, event)
}
