CREATE TABLE db.last_booked
(
    last_user_id INT         NOT NULL,
    table_name   VARCHAR(64) NOT NULL,

    CONSTRAINT last_user_id
        PRIMARY KEY (last_user_id)
)