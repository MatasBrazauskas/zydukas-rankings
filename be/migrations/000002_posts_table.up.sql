CREATE TABLE IF NOT EXISTS posts (
     id           BIGSERIAL PRIMARY KEY,
     user_id      BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
     caption      TEXT NOT NULL DEFAULT '',
     picture_path TEXT NOT NULL UNIQUE,
     created_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_posts_user_id ON posts (user_id);
CREATE INDEX IF NOT EXISTS idx_posts_feed ON posts (created_at DESC, id DESC);
