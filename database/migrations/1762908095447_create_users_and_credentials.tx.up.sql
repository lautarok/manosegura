-- habilita uuid_generate_v4
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Tabla users
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(40) NOT NULL,
    surname VARCHAR(40) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Tabla credentials
CREATE TABLE credentials (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(50) UNIQUE NOT NULL,
    username VARCHAR(26) UNIQUE NOT NULL,
    password VARCHAR(60) NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
