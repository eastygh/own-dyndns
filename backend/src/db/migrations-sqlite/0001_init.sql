-- +goose Up
create table od_user
(
    id       integer not null
        constraint user_pk
            primary key autoincrement,
    login    text    not null,
    password TEXT    not null,
    active   integer default 0
);

create table od_parameter
(
    id        integer not null
        constraint parameters_pk
            primary key autoincrement,
    name      TEXT    not null,
    value_str TEXT,
    value_int integer
);

insert into od_parameter (name, value_int)
values ("init", 0);

-- +goose Down
DROP TABLE od_user;

DROP TABLE od_parameter;
