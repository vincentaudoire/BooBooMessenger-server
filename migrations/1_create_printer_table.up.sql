CREATE TABLE Printer (
  id            SERIAL PRIMARY KEY,
  uuid          INT             NOT NULL UNIQUE,
  name          Text            NOT NULL
)
