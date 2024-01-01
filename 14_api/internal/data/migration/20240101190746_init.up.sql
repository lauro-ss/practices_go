CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create or replace function updateAt()
returns trigger as $$
begin
    new."updateAt" = now();
    return new;
end;
$$ language plpgsql;

CREATE TABLE IF NOT EXISTS animal (
  id UUID DEFAULT uuid_generate_v4 (),
  name VARCHAR(20) not null,
  emoji VARCHAR(10) not null,
  createdAt timestamp not null default now(),
  updatedAt timestamp null,
  primary key (id)
);

create trigger upAnimal
before update on animal
for each row execute procedure updateAt();

create table if not exists food (
    id uuid primary key default uuid_generate_v4 (),
    name varchar(10) not null,
    emoji varchar(10) not null,
    createdAt timestamp not null default now(),
    updatedAt timestamp null
);

create trigger upFood
before update on food
for each row execute procedure updateAt();

create table if not exists animalFood (
    idAnimal uuid not null references animal(id),
    idFood uuid not null references food,

    primary key (idAnimal, idFood)
);