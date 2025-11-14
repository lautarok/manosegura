CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    alias VARCHAR(30) UNIQUE NOT NULL,
    name_en VARCHAR(30),
    name_es VARCHAR(30),
    name_fr VARCHAR(30),
    name_pt VARCHAR(30),
    CHECK(
        name_en NOT NULL OR
        name_es NOT NULL OR
        name_fr NOT NULL OR
        name_pt NOT NULL
    )
);

CREATE TABLE permissions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    alias VARCHAR(50) UNIQUE NOT NULL,
    name_en VARCHAR(30),
    name_es VARCHAR(30),
    name_fr VARCHAR(30),
    name_pt VARCHAR(30),
    CHECK(
        name_en NOT NULL OR
        name_es NOT NULL OR
        name_fr NOT NULL OR
        name_pt NOT NULL
    )
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

INSERT INTO roles (
    alias,
    name_en,
    name_es,
    name_fr,
    name_pt
) VALUES (
    'admin',
    'Administrator',
    'Administrador',
    'Administrateur',
    'Administrador'
);
INSERT INTO permissions (
    alias,
    name_en,
    name_es,
    name_fr,
    name_pt
) VALUES (
    'manage all',
    'Manage all',
    'Administrar todo',
    'Gérer tout',
    'Gerenciar tudo'
);
INSERT INTO role_permissions (permission_id, role_id) VALUES (
    (SELECT id FROM permissions WHERE alias = 'manage all'),
    (SELECT id FROM roles WHERE alias = 'admin')
);

INSERT INTO roles (
    alias,
    name_en,
    name_es,
    name_fr,
    name_pt
) VALUES (
    'regular user',
    'Regular user',
    'Usuario regular',
    'Utilisateur régulier',
    'Usuário normal'
);

INSERT INTO permissions (
    alias,
    name_en,
    name_es,
    name_fr,
    name_pt
) VALUES (
    'regular user',
    'Regular user',
    'Usuario regular',
    'Utilisateur régulier',
    'Usuário normal'
);
INSERT INTO role_permissions (permission_id, role_id) VALUES (
    (SELECT id FROM permissions WHERE alias = 'regular user'),
    (SELECT id FROM roles WHERE alias = 'regular user')
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