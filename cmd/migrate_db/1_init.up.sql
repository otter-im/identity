CREATE TABLE users
(
    id            uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    username      varchar(255) UNIQUE NOT NULL,
    password_hash varchar(255)        NOT NULL,
    creation_date timestamp DEFAULT current_timestamp
);

/* tinyfluffs | changeme */
INSERT into users(username, password_hash)
VALUES ('tinyfluffs', '$2a$12$qqPrblVzy2HqkM6ZBvbNL.UjbTH1WOFDNxjBX2aMsgl3sS1vfJgUG');
