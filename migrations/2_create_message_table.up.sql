CREATE TABLE Message (
  id                 SERIAL PRIMARY KEY,
  text               TEXT           NOT NULL,
  received           TIMESTAMP      NOT NULL,
  printed            TIMESTAMP      NULL,
  printer_id         INT            NOT NULL,
  FOREIGN  KEY (printer_id) REFERENCES printer(id)
)
