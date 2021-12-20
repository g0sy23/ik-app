CREATE TABLE IF NOT EXISTS merch_categories (
    id          BIGSERIAL    PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS merch_items (
    id          BIGSERIAL    PRIMARY KEY,
    category_id BIGSERIAL    REFERENCES merch_categories,
    title       VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    quantity    INTEGER      NOT NULL,
    size        INTEGER      NOT NULL,
    color       INTEGER      NOT NULL
);