CREATE TYPE genre AS ENUM ('technology', 'education', 'philosophy', 'politics', 'business', 'health', 'sports', 'religion');
CREATE TYPE interaction_type AS ENUM ('like', 'listen');

CREATE TABLE podcast_metadata (
    podcast_id uuid not null    ,
    genre genre not null,
    tags TEXT[],
    listen_count INTEGER DEFAULT 0,
    like_count INTEGER DEFAULT 0,
    created_at TIMESTAMP,
    updated_at TIMESTAMP

);

CREATE TABLE user_interactions (
    id uuid PRIMARY KEY not null,
    user_id uuid,
    podcast_id uuid,
    interaction_type interaction_type not null ,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
