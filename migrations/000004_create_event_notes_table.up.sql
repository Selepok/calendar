CREATE TABLE IF NOT EXISTS event_notes
(
    id       SERIAL PRIMARY KEY,
    event_id INT NOT NULL,
    item     VARCHAR(50)
);
