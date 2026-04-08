CREATE TABLE bundle (
    slug TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_accessed TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_bundle_last_accessed ON bundle(last_accessed);

CREATE TABLE link (
    url TEXT NOT NULL,
    note TEXT,
    display_text TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    bundle_slug TEXT NOT NULL REFERENCES bundle(slug) ON DELETE CASCADE,
    PRIMARY KEY (bundle_slug, url)
);