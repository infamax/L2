package service

import "context"

func (s *service) DeleteEvent(ctx context.Context, eventID int) error {
	return s.repo.DeleteEvent(ctx, eventID)
}
