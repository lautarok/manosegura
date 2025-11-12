CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    alias VARCHAR(30) UNIQUE NOT NULL
);

CREATE TABLE permissions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    alias VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE role_permissions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    permission_id UUID NOT NULL REFERENCES permissions(id) ON DELETE CASCADE,
    role_id UUID NOT NULL REFERENCES roles(id) ON DELETE CASCADE
);

ALTER TABLE users
ADD COLUMN role_id UUID;

ALTER TABLE users
ADD CONSTRAINT fk_role_id
FOREIGN KEY (role_id)
REFERENCES roles(id);

INSERT INTO roles (alias) VALUES ('admin');
INSERT INTO permissions (alias) VALUES ('manage all');
INSERT INTO role_permissions (permission_id, role_id) VALUES (
    (SELECT id FROM permissions WHERE alias = 'manage all'),
    (SELECT id FROM roles WHERE alias = 'admin')
);

INSERT INTO users (name, surname, role_id) VALUES (
    'Lautaro',
    'Kazalukian',
    (SELECT id FROM roles WHERE alias = 'admin')
);

INSERT INTO credentials (email, username, password, user_id) VALUES (
    'admin@manosegura.com.ar',
    'admin.manosegura',
    -- Password: $Administrador2020
    '$2y$10$NxuWCFtQShMmuKSz1bo/juA2eYWVhKCjKAhbgZagul5iHtzI2dkF6',
    (SELECT id FROM users WHERE name = 'Lautaro' and surname = 'Kazalukian')
);