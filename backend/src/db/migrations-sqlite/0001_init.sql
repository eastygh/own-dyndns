-- +goose Up
create table od_user
(
    id       integer             not null
        constraint user_pk
            primary key autoincrement,
    login    TEXT collate NOCASE not null,
    password TEXT                not null,
    active   integer default 0,
    UNIQUE (login COLLATE NOCASE)
);

create table od_role
(
    id        integer             not null
        constraint role_pk
            primary key autoincrement,
    role_name TEXT collate NOCASE not null
        constraint role_uk
            unique,
    role_desc TEXT
);

insert into od_role (role_name, role_desc) values ('suPer_admin','Super admin of Own DynDNS application');
insert into od_role (role_name, role_desc) values ('doMain_admin','Admin of some serv domain');
insert into od_role (role_name, role_desc) values ('domain_updater','Role for dynamic domain updater');

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
values ('init', 0);

-- +goose Down
DROP TABLE od_user;

DROP TABLE od_parameter;
