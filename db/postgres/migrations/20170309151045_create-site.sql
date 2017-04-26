-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE hosts (
  id            BIGSERIAL PRIMARY KEY,
  lang          VARCHAR(255)                NOT NULL,
  name          VARCHAR(255)                NOT NULL,
  root          BOOLEAN                     NOT NULL DEFAULT FALSE,
  title         VARCHAR(255)                NOT NULL,
  sub_title     VARCHAR(255)                NOT NULL,
  author_id     BIGINT                      REFERENCES users,
  keywords      VARCHAR(255)                NOT NULL,
  description   VARCHAR(800)                NOT NULL,
  copyright     VARCHAR(255)                NOT NULL,
  ssl           BOOLEAN                     NOT NULL DEFAULT FALSE,
  public_perm   TEXT,
  private_perm  TEXT,
  created_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_hosts_name_lang ON hosts(name, lang);
CREATE INDEX idx_hosts_name ON hosts(name);


CREATE TABLE leave_words (
  id         BIGSERIAL PRIMARY KEY,
  body       TEXT                        NOT NULL,
  type       VARCHAR(8)                  NOT NULL DEFAULT 'markdown',
  host_id    BIGINT                      REFERENCES hosts,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE posts (
  id         BIGSERIAL PRIMARY KEY,
  name        VARCHAR(32)                 NOT NULL,
  lang       VARCHAR(8)                 NOT NULL,
  title      VARCHAR(255)                NOT NULL,
  body       TEXT                        NOT NULL,
  type       VARCHAR(8)                  NOT NULL DEFAULT 'markdown',
  host_id    BIGINT                      REFERENCES hosts,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_posts_name_lang_host ON posts(name, lang, host_id);
CREATE INDEX idx_posts_name ON posts(name);
CREATE INDEX idx_posts_lang ON posts(lang);

CREATE TABLE notices (
  id         BIGSERIAL PRIMARY KEY,
  body       TEXT                        NOT NULL,
  type       VARCHAR(8)                  NOT NULL DEFAULT 'markdown',
  host_id    BIGINT                      REFERENCES hosts,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE links (
  id          BIGSERIAL PRIMARY KEY,
  href        VARCHAR(255) NOT NULL,
  label       VARCHAR(255) NOT NULL,
  loc         VARCHAR(16) NOT NULL,
  sort_order  INT NOT NULL DEFAULT 0,
  host_id     BIGINT                      REFERENCES hosts,
  created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX idx_links_loc ON links (loc);

CREATE TABLE pages (
  id          BIGSERIAL PRIMARY KEY,
  title       VARCHAR(255) NOT NULL,
  summary     VARCHAR(2048) NOT NULL,
  action      VARCHAR(32) NOT NULL,
  href        VARCHAR(255) NOT NULL,
  logo        VARCHAR(255) NOT NULL,
  loc         VARCHAR(16) NOT NULL,
  sort_order  INT NOT NULL DEFAULT 0,
  host_id     BIGINT                      REFERENCES hosts,
  created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX idx_pages_loc ON pages (loc);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE pages;
DROP TABLE links;
DROP TABLE notices;
DROP TABLE posts;
DROP TABLE leave_words;
