drop table if exists users;
drop table if exists tournaments;
drop table if exists participants;

create table if not exists users(
    id serial constraint user_pk primary key,
    name text not null,
    balance int not null
);

create table if not exists tournaments(
    id serial constraint tournament_pk primary key,
    name text not null,
    deposit int not null,
    prize int not null
);

create table if not exists participants(
    id serial constraint participants_pk primary key,
    tournamentid serial,
    userid serial
);

create table if not exists winners(
   id serial constraint winners_pk primary key,
   tournamentid serial,
   userid serial
);