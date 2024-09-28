DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS instances;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL, 
    email VARCHAR(50) NOT NULL UNIQUE, 
    password VARCHAR(255) NOT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);

CREATE TABLE instances (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    ram BIGINT,
    disk VARCHAR(255),
    vcpus BIGINT,
    osvariant VARCHAR(255),
    console VARCHAR(255),
    location VARCHAR(255),
    extraargs TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

