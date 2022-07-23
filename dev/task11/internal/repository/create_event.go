package repository

import (
	"context"
	"github.com/infamax/l2/task11/internal/models"
)

func (d *postgresDB) CreateEvent(ctx context.Context, event *models.Event) error {
	const query = `
		insert into events
		(user_id, title, date_created, date_finished,
		description, done)
		values($1, $2, $3, $4, $5, $6)
		returning event_id;
	`
	err := d.pool.QueryRow(ctx, query, event.UserID, event.Title,
		event.DateCreated, event.DateFinished, event.Description, event.Done).Scan(&event.EventID)

	return err
}
