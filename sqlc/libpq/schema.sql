CREATE TABLE products (
    id          uuid         not null default uuid_generate_v4() primary key,
    name        varchar(100) not null,
    description text         not null,
    categories  text[]       not null,
    price       float        not null,
    features    text[]       not null,
    color       text         not null,
    material    text         not null,
    upc         text         not null
);