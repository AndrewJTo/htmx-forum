CREATE TABLE users (
    id          SERIAL          PRIMARY KEY,
    email       VARCHAR(255)    NOT NULL,
    name        VARCHAR(255)    NOT NULL,
    password    VARCHAR(255)    NOT NULL,
    is_admin    BOOLEAN         DEFAULT false,
    ban_until   DATE            DEFAULT NULL,
    created_at  DATE            NOT NULL
);

CREATE TABLE category (
    id              SERIAL          PRIMARY KEY,
    name            VARCHAR(255)    NOT NULL,
    parent_id       INTEGER         REFERENCES category,
    description     TEXT            DEFAULT NULL,
    image_id        INTEGER         REFERENCES images,
    creator_user_id INTEGER         REFERENCES users,
    created_at      DATE            NOT NULL
    -- FOREIGN KEY (creator_user_id) REFERENCES user (id)
);

CREATE TABLE thread (
    id              SERIAL          PRIMARY KEY,
    name            VARCHAR(255)    NOT NULL,
    category_id     INTEGER         DEFAULT NULL,
    creator_user_id INTEGER         DEFAULT NULL,
    created_at      DATE            NOT NULL
);

CREATE TABLE post (
    id              SERIAL          PRIMARY KEY,
    thread_id       INTEGER         DEFAULT NULL,
    creator_user_id INTEGER         DEFAULT NULL,
    image_id        INTEGER         DEFAULT NULL,
    content         TEXT            DEFAULT NULL,
    created_at      DATE            NOT NULL
);

CREATE TABLE image (
    id              SERIAL          PRIMARY KEY,
    creator_user_id INTEGER         DEFAULT NULL,
    created_at      DATE            NOT NULL
);