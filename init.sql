-- init.sql
CREATE TABLE IF NOT EXISTS stocks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);

INSERT INTO stocks (name, price) VALUES
('Apple', 150.00),
('Google', 2800.00),
('Amazon', 3400.00);