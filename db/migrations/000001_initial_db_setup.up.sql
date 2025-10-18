-- USERS TABLE
CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    name varchar(15) NOT NULL UNIQUE,
    password varchar NOT NULL
);

-- ACCOUNTS TABLE
CREATE TABLE IF NOT EXISTS accounts (
    id serial PRIMARY KEY,
    document_number varchar(20) NOT NULL UNIQUE
);

-- OPERATION TYPES TABLE
CREATE TABLE IF NOT EXISTS operation_types (
    id serial PRIMARY KEY,
    description varchar NOT NULL
);

-- TRANSACTIONS TABLE
CREATE TABLE IF NOT EXISTS transactions (
    id serial PRIMARY KEY,
    account_id int NOT NULL,
    operation_type_id int NOT NULL,
    amount float NOT NULL,
    event_date timestamp DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES accounts(id),
    FOREIGN KEY (operation_type_id) REFERENCES operation_types(id)
);
