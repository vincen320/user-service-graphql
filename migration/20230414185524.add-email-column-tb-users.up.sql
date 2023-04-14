ALTER TABLE users
    ADD COLUMN email varchar

CREATE UNIQUE INDEX ON users (LOWER("email"))

UPDATE users SET
email = 'budiana@mail.com'
WHERE id = 1

UPDATE users SET
email = 'stefy@mail.com'
WHERE id = 2

UPDATE users SET
email = 'lena@mail.com'
WHERE id = 3

UPDATE users SET
email = 'markus@mail.com'
WHERE id = 4

UPDATE users SET
email = 'kevin@mail.com'
WHERE id = 5

UPDATE users SET
email = 'vincen@mail.com'
WHERE id = 6

UPDATE users SET
email = 'giovany@mail.com'
WHERE id = 7