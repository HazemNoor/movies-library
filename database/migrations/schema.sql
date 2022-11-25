create table users
(
    id       int unsigned auto_increment primary key,
    name     varchar(255) not null,
    age      int unsigned not null,
    email    varchar(255) not null,
    password varchar(100) not null,
    constraint users_email_unique
        unique (email)
);

create table movies
(
    id          int unsigned auto_increment primary key,
    name        varchar(255)   not null,
    description text           not null,
    date        date           not null,
    cover       varchar(255)   not null,
    rate        float unsigned not null,
    user_id     int unsigned   not null,
    created_at  datetime       not null,
    updated_at  datetime       not null,
    constraint movies_users_fk
        foreign key (user_id) references users (id)
            on update cascade on delete cascade
);

create index movies_user_id_index on movies (user_id);

create table user_tokens
(
    id         bigint auto_increment primary key,
    user_id    int unsigned not null,
    token      varchar(255) not null,
    created_at datetime     not null,
    expires_at datetime     not null,
    constraint user_tokens_users_fk
        foreign key (user_id) references users (id)
            on update cascade on delete cascade
);

create index user_tokens_user_id_index on user_tokens (user_id);

create table watched_list_items
(
    id         int unsigned auto_increment primary key,
    user_id    int unsigned     not null,
    movie_id   int unsigned     not null,
    rate       tinyint unsigned not null,
    review     text             not null,
    created_at datetime         not null,
    updated_at datetime         not null,
    constraint watched_list_items_movies_fk
        foreign key (movie_id) references movies (id)
            on update cascade on delete cascade,
    constraint watched_list_items_users_fk
        foreign key (user_id) references users (id)
            on update cascade on delete cascade
);
