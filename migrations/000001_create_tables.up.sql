CREATE TYPE genre AS ENUM ('technology', 'education', 'philosophy', 'politics', 'business', 'health', 'sports', 'religion');
CREATE TYPE interaction_type AS ENUM ('like', 'listen');

CREATE TABLE episode_metadata (
    episode_id uuid primary key,
    podcast_id uuid not null,
    genre genre not null,
    tags TEXT[],
    listen_count INTEGER DEFAULT 0,
    like_count INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE user_interactions (
    id uuid PRIMARY KEY default gen_random_uuid(),
    user_id uuid not null,
    podcast_id uuid not null,
    episode_id uuid not null,
    interaction_type interaction_type not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp
);
