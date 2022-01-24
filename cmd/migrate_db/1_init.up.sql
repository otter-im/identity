CREATE TABLE users
(
    id            uuid      DEFAULT gen_random_uuid() PRIMARY KEY,
    username      varchar(255) UNIQUE NOT NULL,
    hash          bytea               NOT NULL,
    salt          bytea               NOT NULL,
    creation_date timestamp DEFAULT current_timestamp
);

/* tinyfluffs | changeme */
INSERT into users(id, username, hash, salt, creation_date)
VALUES ('d6ef6dc7-ce36-449c-8265-07f60ca3b2ff',
        'tinyfluffs',
        decode('436c0dbc0a78026e8ce9dcf8a95cc1a4bcefaee55ee70d87874a27fb0829b994', 'hex'),
        decode('2780cb19d7f864d49179ffb725284fa0', 'hex'),
        '2022-01-24 12:00:50.884460');
