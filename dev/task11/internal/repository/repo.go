package repository

import (
	"context"
	"github.com/infamax/l2/task11/internal/models"
)

type Repository interface {
	CreateEvent(context.Context, *models.Event) error
	GetEvents(context.Context, int) ([]models.Event, error)
	UpdateEvent(context.Context, *models.Event) error
	DeleteEvent(context.Context, int) error
}
