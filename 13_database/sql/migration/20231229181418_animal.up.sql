CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS animal (
  id UUID DEFAULT uuid_generate_v4 (),
  name VARCHAR(20) not null,
  emoji VARCHAR(10) not null,
  primary key (id)
);