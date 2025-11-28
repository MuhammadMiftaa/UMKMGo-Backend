-- config/db/migrations/20251128120000_create_news.sql

-- +goose Up
-- +goose StatementBegin
CREATE TYPE news_category AS ENUM (
    'announcement',
    'event',
    'program_update',
    'success_story',
    'tips',
    'regulation',
    'general'
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE news (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    excerpt TEXT,
    content TEXT NOT NULL,
    thumbnail TEXT,
    category news_category NOT NULL DEFAULT 'general',
    author_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    is_published BOOLEAN NOT NULL DEFAULT false,
    published_at TIMESTAMPTZ,
    views_count INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX idx_news_slug ON news(slug);
CREATE INDEX idx_news_category ON news(category);
CREATE INDEX idx_news_published ON news(is_published, published_at DESC);
CREATE INDEX idx_news_author ON news(author_id);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE news_tags (
    id SERIAL PRIMARY KEY,
    news_id INT NOT NULL REFERENCES news(id) ON DELETE CASCADE,
    tag_name VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE(news_id, tag_name)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX idx_news_tags_news_id ON news_tags(news_id);
CREATE INDEX idx_news_tags_tag_name ON news_tags(tag_name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS news_tags;
DROP TABLE IF EXISTS news;
DROP TYPE IF EXISTS news_category;
-- +goose StatementEnd