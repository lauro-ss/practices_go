CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS animal (
  id UUID primary key,
  name VARCHAR(20) not null,
  emoji VARCHAR(10)
);