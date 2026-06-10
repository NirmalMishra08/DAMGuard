


CREATE TABLE  users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
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