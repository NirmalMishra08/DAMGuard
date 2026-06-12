

CREATE TYPE user_role AS ENUM ('USER', 'TRAINER', 'ADMIN');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    phone TEXT,
    email TEXT UNIQUE NOT NULL,
    fullname TEXT NOT NULL,
    provider provider NOT NULL,
    password_hash TEXT,
    role user_role NOT NULL DEFAULT 'USER',
    updated_at timestamptz DEFAULT NOW()
);


CREATE TABLE database_sources (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    host VARCHAR(255) NOT NULL,
    port INT NOT NULL,
    database_name VARCHAR(255),
    username VARCHAR(255),
    password_encrypted TEXT,
    enabled BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW()
);


CREATE TABLE alerts (
    id UUID PRIMARY KEY,
    database_id UUID REFERENCES database_sources(id),
    severity VARCHAR(20),
    title VARCHAR(255),
    description TEXT,
    status VARCHAR(20) DEFAULT 'open',
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE alert_history (
    id UUID PRIMARY KEY,
    alert_id UUID,
    action VARCHAR(50),
    performed_by UUID,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE settings (
    key VARCHAR(255) PRIMARY KEY,
    value TEXT
);