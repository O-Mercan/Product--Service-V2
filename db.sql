CREATE DATABASE asd;


CREATE TABLE products(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    category VARCHAR(100) NOT NULL,
    summary VARCHAR(100) NOT NULL,
    description VARCHAR(150) NOT NULL,
    price int NOT NULL
);
