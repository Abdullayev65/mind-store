CREATE TABLE IF NOT EXISTS users
(
    id          BIGSERIAL PRIMARY KEY,
    username    VARCHAR(26)             NOT NULL,
    email       VARCHAR(40),
    mind_id     BIGINT UNIQUE,
    password    VARCHAR(30)             NOT NULL,
    first_name  VARCHAR(16),
    middle_name VARCHAR(16),
    last_name   VARCHAR(20),
    birth_date  TIMESTAMP,
    avatar_id   BIGINT,
--
    created_by  BIGINT,
    deleted_by  BIGINT,
    created_at  TIMESTAMP DEFAULT now() NOT NULL,
    deleted_at  TIMESTAMP
);
CREATE UNIQUE INDEX users_username_unique_index
    ON users (username)
    WHERE users.deleted_at IS NULL;
CREATE UNIQUE INDEX users_email_unique_index
    ON users (email)
    WHERE users.deleted_at IS NULL;

CREATE TABLE IF NOT EXISTS mind
(
    id         BIGSERIAL PRIMARY KEY,
    topic      VARCHAR(40) NOT NULL,
    caption    TEXT,
    parent_id  BIGINT REFERENCES mind (id),
    access     BIGINT      NOT NULL,
    hashed_id  BIGINT REFERENCES mind (id),
--
    created_at TIMESTAMP   NOT NULL DEFAULT now(),
    updated_at TIMESTAMP   NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP,
    created_by BIGINT REFERENCES users (id),
    deleted_by BIGINT REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS file
(
    id         BIGSERIAL PRIMARY KEY,
    name       VARCHAR(63)  NOT NULL,
    path       VARCHAR(127) NOT NULL,
    hashed_id  BIGINT REFERENCES mind (id),
    access     BIGINT       NOT NULL,
    size       BIGINT       NOT NULL,
--
    created_at TIMESTAMP    NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP,
    created_by BIGINT REFERENCES users (id),
    deleted_by BIGINT REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS mind_file
(
    mind_id    BIGINT    NOT NULL REFERENCES mind (id),
    file_id    BIGINT    NOT NULL REFERENCES file (id),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP,
    PRIMARY KEY (mind_id, file_id)
);
CREATE UNIQUE INDEX mind_file_ui__mind_id__file_id ON mind_file
    USING btree (mind_id, file_id) WHERE (deleted_at IS NULL);


ALTER TABLE users
    ADD FOREIGN KEY (mind_id) REFERENCES mind (id),
    ADD FOREIGN KEY (avatar_id) REFERENCES file (id);

