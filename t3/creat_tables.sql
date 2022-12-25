create database db1;
use db1;
create table user
(
    uid      int primary key auto_increment,
    sex      varchar(1)  not null,
    username varchar(20) not null unique,
    password varchar(20) not null,
    phone    varchar(11) not null,
    email    varchar(50) not null
);
create table log
(
    lid         int primary key auto_increment,
    l_uid       int         not null,
    borrow_at   time        not null,
    last_time   int         not null,
    l_bid       varchar(30) not null,
    is_return   bool,
    return_at   time,
    is_overtime bool,
    overtime    int
);
create table books
(
    bid           varchar(30) primary key,
    name          varchar(50) not null,
    author        varchar(50) not null,
    publish_place varchar(30) not null,
    publish_at    time        not null,
    library       int         not null,
    total_number  int         not null,
    borrow_number int         not null,
    remain_number int         not null
);