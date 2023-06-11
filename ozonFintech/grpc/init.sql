create database ozonlinksdb;
GRANT ALL PRIVILEGES ON DATABASE ozonlinksdb TO postgres;
\c ozonlinksdb;
create table if not exists links(
    abbreviated varchar(10) primary key,
    original varchar(256)
);