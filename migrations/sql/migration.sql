CREATE TABLE test.users (
		id varchar(36) DEFAULT (uuid()) NOT NULL,
		email varchar(255) UNIQUE NOT NULL,
		password varchar(255) NOT NULL,
		balance int NOT NULL DEFAULT 0,
		created_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
		updated_at TIMESTAMP DEFAULT current_timestamp ON UPDATE current_timestamp NOT NULL,
		CONSTRAINT users_PK PRIMARY KEY (id)
	);

CREATE TABLE test.products (
		id varchar(36) DEFAULT (uuid()) NOT NULL,
		name varchar(255) NOT NULL,
		description text,
		category varchar(255) NOT NULL,
		price int NOT NULL DEFAULT 0,
		created_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
		updated_at TIMESTAMP DEFAULT current_timestamp ON UPDATE current_timestamp NOT NULL,
		CONSTRAINT products_PK PRIMARY KEY (id)
	);

CREATE TABLE test.carts (
		id varchar(36) DEFAULT (uuid()) NOT NULL,
		user_id varchar(255) NOT NULL REFERENCES users(id),
		product_id varchar(255) NOT NULL REFERENCES products(id),
		amount int NOT NULL DEFAULT 1,
		created_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
		updated_at TIMESTAMP DEFAULT current_timestamp ON UPDATE current_timestamp NOT NULL,
		CONSTRAINT carts_PK PRIMARY KEY (id)
	);

CREATE TABLE test.payments (
		id varchar(36) DEFAULT (uuid()) NOT NULL,
		cart_id varchar(255) NOT NULL REFERENCES carts(id),
		created_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
		updated_at TIMESTAMP DEFAULT current_timestamp ON UPDATE current_timestamp NOT NULL,
		CONSTRAINT carts_PK PRIMARY KEY (id)
	);

INSERT INTO products(name, description, category, price) VALUES
		('Mango Juice', 'Fresh', 'Drink', 10000),
		('Fried Tofu', 'Delicious', 'Food', 1000);