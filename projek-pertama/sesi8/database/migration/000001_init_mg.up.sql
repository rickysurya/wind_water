CREATE TABLE employees (
    id serial primary key,
    full_name varchar(50) not null,
    email text unique not null,
    age int not null,
    division varchar(20) not null
);