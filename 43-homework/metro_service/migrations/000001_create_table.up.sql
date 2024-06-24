CREATE TABLE card (
    id UUID PRIMARY KEY,
    number VARCHAR(16) NOT NULL,
    user_id UUID,
    FOREIGN KEY (user_id) REFERENCES user2(id) ON DELETE CASCADE
);

create table station 
(
    id uuid PRIMARY KEY,
    name varchar(20)
);

CREATE TABLE terminal
(
    id uuid PRIMARY KEY,
    station_id uuid,
    FOREIGN KEY (station_id) REFERENCES station(id) ON DELETE CASCADE
);

CREATE TYPE transaction_type_enum AS ENUM ('credit', 'debit');

CREATE TABLE transaction
(
    id uuid PRIMARY KEY,
    card_id uuid,
    amount integer,
    terminal_id uuid,
    transaction_type transaction_type_enum,
    FOREIGN KEY (card_id) REFERENCES card(id) ON DELETE CASCADE,
    FOREIGN KEY (terminal_id) REFERENCES terminal(id) ON DELETE CASCADE
);
