CREATE TABLE tables
(
    id            INT         NOT NULL AUTO_INCREMENT,
    table_name    VARCHAR(64) NOT NULL,
    restaurant_id INT         NOT NULL,

    CONSTRAINT id
        PRIMARY KEY (id),

    CONSTRAINT restaurant_id
        FOREIGN KEY (restaurant_id) REFERENCES db.restaurants (id)
            ON DELETE CASCADE
            ON UPDATE RESTRICT
);