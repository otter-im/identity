CREATE TABLE users
(
    id            uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    username      varchar(255) UNIQUE NOT NULL,
    password_hash varchar(255)        NOT NULL,
    creation_date timestamp DEFAULT current_timestamp
);

/* tinyfluffs | changeme */
INSERT into users(id, username, password_hash, creation_date)
VALUES ('d6ef6dc7-ce36-449c-8265-07f60ca3b2ff', 'tinyfluffs', '$2a$12$qqPrblVzy2HqkM6ZBvbNL.UjbTH1WOFDNxjBX2aMsgl3sS1vfJgUG', '2022-01-24 12:00:50.884460');
