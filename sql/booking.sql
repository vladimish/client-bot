CREATE TABLE booking
(
    id       INT NOT NULL AUTO_INCREMENT,
    table_id INT NOT NULL,
    start    TIMESTAMP,
    end      TIMESTAMP,

    CONSTRAINT id
        PRIMARY KEY (id)
);