CREATE TABLE IF NOT EXISTS events
(
    id
                SERIAL
        PRIMARY
            KEY,
    user_id
                INT
        NOT
            NULL,
    title
                VARCHAR(100),
    description text,
    time        TIMESTAMP,
    timezone    VARCHAR(50),
    duration    INT
);
