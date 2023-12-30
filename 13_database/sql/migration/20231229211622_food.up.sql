create table if not exists food (
    id uuid primary key default uuid_generate_v4 (),
    name varchar(10) not null,
    emoji varchar(10) not null
);