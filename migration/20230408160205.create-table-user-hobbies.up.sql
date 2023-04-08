CREATE TABLE user_hobbies(
    id bigserial PRIMARY KEY
    , user_id bigint REFERENCES users(id)
    , hobby_id bigint REFERENCES hobbies(id)
)

INSERT INTO user_hobbies
(user_id, hobby_id)
VAlUES
(1, 1),
(1, 2),
(1, 3),
(2, 6),
(2, 8),
(2, 10),
(3, 16),
(1, 17),
(3, 19),
(4, 1),
(4, 22),
(4, 26),
(2, 3),
(3, 10),
(5, 17)