use kh;
create table users(
    uid int primary key auto_increment,
    username varchar(20) unique not null ,
    password varchar(20) not null
)