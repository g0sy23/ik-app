CREATE TABLE merch_categories (
  id           INTEGER      GENERATED ALWAYS AS IDENTITY,
  title        VARCHAR(20)  NOT NULL,
  description  VARCHAR(40),
  CONSTRAINT merch_categories_pk
    PRIMARY KEY (id)
);

INSERT INTO merch_categories (title, description) VALUES ('merch', 'default category');

CREATE RULE protect_first_entry_update AS ON update
  TO merch_categories WHERE old.id = 1
  DO INSTEAD NOTHING;

CREATE RULE protect_first_entry_delete AS ON delete
  TO merch_categories WHERE old.id = 1
  DO INSTEAD NOTHING;

CREATE TABLE merch_items (
  id           INTEGER      GENERATED ALWAYS AS IDENTITY,
  category_id  INTEGER      DEFAULT 1,
  title        VARCHAR(20)  NOT NULL,
  description  VARCHAR(40),
  price        INTEGER      NOT NULL,
  quantity     INTEGER      NOT NULL,
  size         VARCHAR(20),
  color        VARCHAR(20),
  CONSTRAINT merch_items_pk
    PRIMARY KEY (id),
  CONSTRAINT merch_items_fk
    FOREIGN KEY (category_id) REFERENCES merch_categories (id)
      ON DELETE SET DEFAULT
);