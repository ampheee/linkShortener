create database ozonlinksdb;
grant all privileges on database ozonlinksdb to postgres;
\c ozonlinksdb;
create table if not exists links(
    abbreviated varchar(10) primary key,
    original varchar(256)
);