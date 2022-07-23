package repository

import (
	"context"
	"github.com/infamax/l2/task11/internal/models"
)

func (d *postgresDB) GetEvents(ctx context.Context, userID int, date string, days int) ([]models.Event, error) {
	const query = `
		select event_id, title, date_created, date_finished,
		description, done
		from events
		where user_id = $1 and
		date_finished >= to_timestamp($2, 'YYYY-MM-DD HH:MI:SS') and 
		date_finished <= to_timestamp($2, 'YYYY-MM-DD HH:MI:SS') + 
		make_interval(days => $3);
	`

	rows, err := d.pool.Query(ctx, query, userID, date, days)

	if err != nil {
		return nil, err
	}

	events := make([]models.Event, 0)

	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.EventID, &event.Title, &event.DateCreated,
			&event.DateFinished, &event.Description, &event.Done)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
