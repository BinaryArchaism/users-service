CREATE TABLE users
(
    id         serial       not null unique,
    first_name varchar(70)  not null,
    last_name  varchar(70)  not null,
    email      varchar(100) not null,
    age        int default null
);