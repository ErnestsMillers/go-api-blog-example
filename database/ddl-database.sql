create table if not exists post (
    post_id bigint not null generated by default as identity,
    post_name varchar not null,
    post_text text not null,
    category_id int not null,

    constraint post_post_id_pkey primary key (post_id),
    constraint post_post_name_ukey unique (post_name),
    constraint post_post_name_check check (post_name <> '')
);

create table if not exists category (
    category_id bigint not null generated by default as identity,
    category_name varchar not null,

    constraint category_category_id_pkey primary key (category_id),
    constraint category_category_name_ukey unique (category_name),
    constraint category_category_name_check check (category_name <> '')
);