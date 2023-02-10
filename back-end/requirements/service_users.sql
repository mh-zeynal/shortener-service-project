create table service_users
(
    user_id    int auto_increment
        primary key,
    username   char(50)                            not null,
    password   text                                not null,
    email      varchar(256)                        not null,
    name       text                                not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    constraint service_users_email_uindex
        unique (email),
    constraint service_users_username_uindex
        unique (username)
);