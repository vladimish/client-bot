CREATE TABLE db.tables
(
    id         INT         NOT NULL AUTO_INCREMENT,
    table_name VARCHAR(64) NOT NULL UNIQUE,

    CONSTRAINT id
        PRIMARY KEY (id)
);