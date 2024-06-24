create table  if not exists station(
    id uuid primary key  default gen_random_uuid(),
    name varchar
);