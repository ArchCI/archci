create database arch;

CREATE TABLE account (
    id integer PRIMARY KEY,
    username varchar(128),
    password varchar(128),
    email varchar(128)
);

insert into account values(1, 'tobe', 'mypass', 'tobe@gmail.com');

insert into account values(2, 'tobe2', 'mypass', 'tobe2@gmail.com');



