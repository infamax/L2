package repository

import (
	"context"
	"github.com/infamax/l2/task11/internal/models"
)

func (d *postgresDB) UpdateEvent(ctx context.Context, event *models.Event) error {
	const query = `
		update events
		set title = $3,
		date_created = $4,
		date_finished = $5,
		description = $6,
		done = $7
		where user_id = $1 and event_id = $2;
	`

	_, err := d.pool.Exec(ctx, query, event.UserID, event.EventID,
		event.Title, event.DateCreated, event.DateFinished, event.Description, event.Done)

	if err != nil {
		return err
	}
	return nil
}
