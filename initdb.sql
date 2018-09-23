CREATE DATABASE phonebook;

USE phonebook;

CREATE TABLE users(
    id bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL,
    surname varchar(100) NOT NULL,
    age integer NOT NULL
);

CREATE TABLE phones(
    id bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    phone varchar(11) NOT NULL,
    user_id bigint NOT NULL,
    FOREIGN KEY(user_id) 
        REFERENCES users(id)
        ON UPDATE CASCADE ON DELETE CASCADE
);