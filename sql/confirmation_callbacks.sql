CREATE TABLE db.confirmation_callbacks
(
    id       INT    NOT NULL AUTO_INCREMENT,
    user_id  BIGINT NOT NULL,
    table_id INT,

    CONSTRAINT id
        PRIMARY KEY (id),

    CONSTRAINT table_id
        FOREIGN KEY (table_id) REFERENCES db.tables (id)
            ON DELETE CASCADE
            ON UPDATE RESTRICT
);