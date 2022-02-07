CREATE TABLE users
(
    id            uuid      DEFAULT gen_random_uuid() PRIMARY KEY,
    username      varchar(255) UNIQUE NOT NULL,
    hash          bytea               NOT NULL,
    salt          bytea               NOT NULL,
    creation_date timestamp DEFAULT current_timestamp
);
GRANT SELECT,INSERT,UPDATE,DELETE ON users TO otter_identity;

/* tinyfluffs | changeme */
INSERT into users(id, username, hash, salt, creation_date)
VALUES ('d6ef6dc7-ce36-449c-8265-07f60ca3b2ff',
        'tinyfluffs',
        decode('70b7a5a7dab303fbf3880c5f75943b53bb1e2818ba9d2330555a78d30e69afd1', 'hex'),
        decode('2780cb19d7f864d49179ffb725284fa0', 'hex'),
        '2022-01-24 12:00:50.884460');
