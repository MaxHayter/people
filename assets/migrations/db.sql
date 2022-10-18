create table person
(
    id         uuid         not null,
    login      varchar(128) not null,
    password   bytea,
    created_at timestamp(0),
    deleted_at timestamp(0),
    primary key (id)
);
create unique index person_index on person using btree (login) where deleted_at is null;
