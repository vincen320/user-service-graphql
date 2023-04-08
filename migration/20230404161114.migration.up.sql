CREATE TABLE IF NOT EXISTS users(
    "id" bigserial PRIMARY KEY,
    "name" varchar,
    "age" int,
    "address" varchar,
    "salary" numeric
);

CREATE INDEX ON users(name)

CREATE INDEX ON users(age)

CREATE INDEX ON users(salary)