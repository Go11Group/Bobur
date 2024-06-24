
create table  if not exists terminal(
                                         id uuid primary key  default gen_random_uuid(),
                                         station_id uuid not null  references station(id)
);