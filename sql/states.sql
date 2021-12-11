CREATE TABLE db.states
(
    user_id    BIGINT NOT NULL,
    state      varchar(8),
    state_data varchar(64),

    CONSTRAINT user_id
        PRIMARY KEY (user_id)
);