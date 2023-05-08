# **How to Use**

## **Create a new Migration**
> ./bin/app migration new -n="filename"

| Flags | Type | Description |
| :--- | :--- | :--- |
| `n` or `name` | `string` | **Optional**. Name of the file|

It will generate 2 files has suffix `up` and `down`\
- `up` is for make changes the database.
- `down` is for reverse the operations performed by the up method.

## **Run the migration**
> ./bin/app migration run {type} {times}

| Args | Type | Description |
| :--- | :--- | :--- |
| `type` | `string` | **Required**. Do `up`/`down` type migration|
| `times` | `int` | **Optional**. How many times migration should be do, if not specified it will migrate all files from the latest migration|

## **Test the Migration**
Note: We are using PostgreSQL as example, maybe you need a little adjustment for the query
1. Run this command
> ./bin/app migration new
2. From generated files that has suffix `up`, write this query:
```sql
CREATE TABLE IF NOT EXISTS users(
    "id" bigserial PRIMARY KEY,
    "name" varchar,
    "age" int,
    "address" varchar,
    "salary" numeric
);
```
3. From generated files that has suffix `down`, write this query:
```sql
DROP TABLE IF EXISTS "users"
```
4. Try Run the Migration
> ./bin/app migration up
5. Check Your Database
6. Maybe you forgot 1 column or has typo on write your SQL, this where `down` operation is used, you can cancel/reverse the operation with `down` migration type
> ./bin/app migration down 1

Please specify how many times the down migration should be do, or it will cancel all migration you has did.