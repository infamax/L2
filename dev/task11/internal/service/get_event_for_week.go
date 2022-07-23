package service

import (
	"context"
	"github.com/infamax/l2/task11/internal/models"
)

func (s *service) GetEventsForWeek(ctx context.Context, userID int, date string) ([]models.Event, error) {
	return s.repo.GetEvents(ctx, userID, date, 7)
}
