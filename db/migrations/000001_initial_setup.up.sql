CREATE TYPE company_type AS ENUM ('Corporations', 'NonProfit', 'Cooperative', 'Sole Proprietorship)');

CREATE TABLE IF NOT EXISTS company (
	id varchar(41) NOT NULL,
	name varchar(15) NOT NULL UNIQUE, 
	description varchar(3000),
	employee_count INT NOT NULL,
	is_registered BOOLEAN NOT NULL,
	type company_type NOT NULL, 
	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
	id varchar(41) NOT NULL,
	name varchar(15) NOT NULL UNIQUE,
	password varchar NOT NULL,
	PRIMARY KEY (id)
);