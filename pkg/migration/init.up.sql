CREATE TABLE IF NOT EXISTS users
(
    id          BIGSERIAL PRIMARY KEY,
    username    VARCHAR(26) UNIQUE      NOT NULL,
    email       VARCHAR(40) UNIQUE,
    mind_id     BIGINT UNIQUE,
    password    VARCHAR(30)             NOT NULL,
    first_name  VARCHAR(16),
    middle_name VARCHAR(16),
    last_name   VARCHAR(20),
    birth_date  TIMESTAMP,
--
    created_by  BIGINT,
    deleted_by  BIGINT,
    created_at  TIMESTAMP DEFAULT now() NOT NULL,
    deleted_at  TIMESTAMP
);

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
    name       VARCHAR(30) NOT NULL,
    path       VARCHAR(50) NOT NULL,
    hashed_id  BIGINT REFERENCES mind (id),
    access     BIGINT      NOT NULL,
    size       BIGINT      NOT NULL,
--
    created_at TIMESTAMP   NOT NULL DEFAULT now(),
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

ALTER TABLE users
    ADD FOREIGN KEY (mind_id) REFERENCES mind (id)
