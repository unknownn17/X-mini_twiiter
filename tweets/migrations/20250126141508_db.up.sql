create table if not exists tweets(
    id serial primary key,
    username varchar(255),
    title varchar(255),
    body text
);
