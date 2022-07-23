CREATE TABLE events(
    user_id int NOT NULL,
    event_id serial NOT NULL,
    title text,
    date_created timestamp DEFAULT NOW(),
    date_finished timestamp,
    description text,
    done bool,
    PRIMARY KEY(user_id, event_id)
)