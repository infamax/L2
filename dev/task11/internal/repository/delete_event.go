package repository

import "context"

func (d *postgresDB) DeleteEvent(ctx context.Context, eventID int) error {
	const query = `
		delete from events
		where event_id = $1;
	`

	_, err := d.pool.Exec(ctx, query, eventID)
	if err != nil {
		return err
	}
	return nil
}
