-- Active: 1719348603208@@127.0.0.1@3306
CREATE TABLE user (
    id          INTEGER         PRIMARY KEY         AUTOINCREMENT,
    email       VARCHAR(255)    NOT NULL,
    name        VARCHAR(255)    NOT NULL,
    password    VARCHAR(255)    NOT NULL,
    is_admin    BOOLEAN         DEFAULT false,
    ban_until   DATE            DEFAULT NULL,
    created_at  DATE            NOT NULL
);

CREATE TABLE category (
    id              INTEGER         AUTOINCREMENT,
    name            VARCHAR(255)    NOT NULL,
    parent_id       INTEGER         DEFAULT NULL,
    description     TEXT            DEFAULT NULL,
    image_id        INTEGER         DEFAULT NULL,
    creator_user_id INTEGER         DEFAULT NULL,
    created_at      DATE            NOT NULL
    PRIMARY KEY (id),
    FOREIGN KEY (creator_user_id) REFERENCES user (id)
);

CREATE TABLE thread (
    id              INTEGER         PRIMARY KEY         AUTOINCREMENT,
    name            VARCHAR(255)    NOT NULL,
    category_id     INTEGER         DEFAULT NULL,
    creator_user_id INTEGER         DEFAULT NULL,
    created_at      DATE            NOT NULL
);

CREATE TABLE post (
    id              INTEGER         PRIMARY KEY         AUTOINCREMENT,
    thread_id       INTEGER         DEFAULT NULL,
    creator_user_id INTEGER         DEFAULT NULL,
    image_id        INTEGER         DEFAULT NULL,
    content         TEXT            DEFAULT NULL,
    created_at      DATE            NOT NULL
);

CREATE TABLE image (
    id              INTEGER         PRIMARY KEY         AUTOINCREMENT,
    creator_user_id INTEGER         DEFAULT NULL,
    created_at      DATE            NOT NULL
);