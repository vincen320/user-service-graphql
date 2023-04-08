CREATE TABLE IF NOT EXISTS hobbies(
    id bigserial PRIMARY KEY
    , name varchar
)

CREATE UNIQUE INDEX ON hobbies (LOWER("name"))

INSERT INTO hobbies(name)
VALUES ('Shopping'), ('Football'), ('Listening to Music'), ('Coding'), ('Fashion'), ('Travelling'), ('Basketball'), ('Gym'), ('Cooking'), ('Study'), ('Gaming'),
       ('Vlogging'), ('Photography'), ('Videography'), ('Art'), ('Playing Music Instrument'), ('Baking'), ('Movies'), ('Volleyball'), ('Dance'), ('Hiking'),
        ('Dogs'), ('Cats'), ('Birds'), ('Fish'), ('Badminton'), ('Running'), ('Jogging'), ('Cycling')