CREATE TABLE Message (
  id                 VARCHAR(255)   NOT NULL,
  message            VARCHAR(255)   NOT NULL,
  received           TIMESTAMP      NOT NULL,
  printed            TIMESTAMP      NULL,
  PRIMARY KEY (id)
)
