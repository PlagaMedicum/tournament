create extension if not exists "uuid-ossp";
drop table if exists users;
drop table if exists tournaments;
drop table if exists participants;
create table users(
    id uuid constraint user_pk primary key default uuid_generate_v4() not null,
    name text not null,
    balance int not null
);
create table tournaments(
    id uuid constraint tournament_pk primary key default uuid_generate_v4() not null,
    name text not null,
    deposit int not null,
    prize int not null,
    users uuid,
    winner uuid
);
create table participants(
     id uuid constraint participant_pk primary key default uuid_generate_v4() not null,
     userid uuid not null
);
insert into users (id, name, balance) values
('bef80618-779e-4cbd-b776-cbd27386a902', 'Samuel Plaunik', 1200);
insert into tournaments (id, name, deposit, prize) values
('6bfccaa8-9e88-4401-a12e-6559e709ee17', 'tour_1', 1000, 4000);