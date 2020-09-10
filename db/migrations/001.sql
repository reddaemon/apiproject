CREATE TABLE rss
(
    id serial not null,
    title varchar,
    link varchar,
    date timestamp unique
);