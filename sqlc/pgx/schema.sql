CREATE TABLE products (
    id          uuid         NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    name        varchar(100) not null,
    price       bigint       not null,
    description text         not null,
    weight      integer      null
);