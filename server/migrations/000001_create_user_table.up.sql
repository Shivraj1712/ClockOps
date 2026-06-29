CREATE TABLE USERS (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email TEXT NOT NULL UNIQUE,
    hash_password TEXT,
    avatar_url TEXT DEFAULT '',
    avatar_public_id TEXT,
    auth_provider TEXT NOT NULL DEFAULT 'local',
    auth_provider_id TEXT,
    user_role VARCHAR(30) NOT NULL DEFAULT 'employee',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() 
)