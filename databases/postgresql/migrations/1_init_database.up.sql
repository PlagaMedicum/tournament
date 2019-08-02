create table if not exists users(
    id serial constraint user_pk primary key,
    name text not null,
    balance int not null
);

create table if not exists tournaments(
    id serial constraint tournament_pk primary key,
    name text not null,
    deposit int not null,
    prize int not null,
    users int,
    winner int
);

create table if not exists participants(
    id serial,
    userid serial
);

insert into users (name, balance) values
('Samuel Plaunik', 1200);
insert into tournaments (name, deposit, prize) values
('tour_1', 1000, 4000);