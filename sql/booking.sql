CREATE TABLE db.booking
(
    id       INT NOT NULL AUTO_INCREMENT,
    table_id INT NOT NULL,
    start    TIMESTAMP,
    end      TIMESTAMP,

    CONSTRAINT id
        PRIMARY KEY (id),

    CONSTRAINT table_id
        FOREIGN KEY (table_id) REFERENCES db.tables (id)
            ON DELETE CASCADE
            ON UPDATE CASCADE
);