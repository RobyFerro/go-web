/* MIGRATION UP */
CREATE TABLE users(
    id bigint unsigned auto_increment PRIMARY KEY ,
    name VARCHAR(255),
    surname VARCHAR(255),
    username VARCHAR(255),
    password VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
)