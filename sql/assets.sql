CREATE TABLE assets (
        id                      SERIAL          PRIMARY KEY     NOT NULL,
        name                    VARCHAR(100)    NOT NULL,
        category                VARCHAR(20)     NOT NULL,
        kind                    VARCHAR(20)     NOT NULL,
        model                   VARCHAR(20)     NOT NULL,
        serial                  VARCHAR(20)     UNIQUE     NOT NULL,
        brand                   VARCHAR(20)     NOT NULL,
        manufactured_year       INT             NOT NULL,
        purchased_date          DATE,
        price                   FLOAT,
        status                  VARCHAR(10)     NOT NULL
);