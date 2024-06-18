CREATE TABLE products (
    id          uuid         default uuid_generate_v4() primary key,
    name        varchar(100) not null,
    description text         not null,
    categories  text[]       not null,
    price       float        not null,
    features    text[]       not null,
    color       text         not null,
    material    text         not null,
    upc         text         not null
);

CREATE TYPE products_v1 AS (
    id          uuid,
    name        varchar(100),
    description text,
    categories  text[],
    price       float,
    features    text[],
    color       text,
    material    text,
    upc         text
);