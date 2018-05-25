create table users(
    id int(6) UNSIGNED AUTO_INCREMENT primary key,
    username varchar(50) not null,
    Pass varchar(50) not null,
    email varchar(50),
    create_date timestamp default CURRENT_TIMESTAMP
);