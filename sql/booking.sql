CREATE TABLE db.booking
(
    id            INT         NOT NULL AUTO_INCREMENT,
    booking_table VARCHAR(64) NOT NULL,
    start         BIGINT,
    end           BIGINT,
    user_id       BIGINT,

    CONSTRAINT id
        PRIMARY KEY (id)
);