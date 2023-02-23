-- auto-generated definition
create table urls
(
    id           int auto_increment
        primary key,
    short_url    text                                not null,
    original_url text                                not null,
    title        text                                null,
    created_at   timestamp default CURRENT_TIMESTAMP null,
    user_id      int                                 null,
    description  text                                null,
    constraint user_id
        foreign key (user_id) references service_users (user_id)
);

