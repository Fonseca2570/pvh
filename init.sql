CREATE DATABASE IF NOT EXISTS sword;
USE sword;

CREATE TABLE migrations
(
    name       varchar(50),
    created_at timestamp DEFAULT CURRENT_TIMESTAMP
);
