# 👋 Mano Segura — Backend

**Mano Segura** es el backend de una plataforma desarrollada en Go que gestiona usuarios, credenciales, roles y permisos. Está diseñado con una arquitectura limpia, soporte para migraciones SQL, documentación Swagger y hot reload en desarrollo.

### 🚀 Tecnologías

-   **Go** 1.25+
    
-   **PostgreSQL**
    
-   **Bun ORM**
    
-   **Air** (hot reload)
    
-   **Swaggo** (Swagger para documentación)
    
-   **uuid-ossp** (para generación de UUIDs en PostgreSQL)
    
-   **JWT**
    
-   **Bcrypt**
    

### 📦 Estructura del proyecto

```
.
├── database
│   └── migrations
├── src
│   ├── docs
│   ├── cmd
│   │   ├── database
│   │   ├── api
│   ├── internal
│   │   ├── exceptions
│   │   ├── modules
│   │   │   ├── [example module]
│   │   │   │   ├── domain
│   │   │   │   ├── dto
│   │   │   │   ├── controllers
│   │   │   │   ├── services
│   │   │   │   ├── repositories
│   │   │   │   ├── interfaces
│   │   │   │   ├── middlewares
│   ├── infra
│   │   ├── factory
│   │   ├── http
│   │   ├── env
│   │   ├── database
└── └── service
```

### ⚙️ Configuración

El backend se configura mediante variables de entorno. Podés definirlas en un archivo .env:

```
HTTP_PORT=3000
POSTGRES_DSN=postgres://user:pass@host:port/database
```

### 🧪 Desarrollo con Air

Instalá Air para hot reload:

bash
`go install github.com/cosmtrek/air@latest`

Luego ejecutá:

```
air
```

Esto recarga automáticamente el servidor al detectar cambios.

### 🛠️ Migraciones

El proyecto incluye una CLI para generar migraciones SQL:

```
go run src/cmd/database/main.go migration create nombre_de_la_migracion
```

Esto crea archivos .tx.up.sql y .tx.down.sql en database/migrations.

### 📚 Documentación Swagger

La API está documentada con Swaggo. Para generar la documentación:

```
swag init -g src/cmd/main.go -o docs --parseDependency
```

Luego accedé a http://localhost:3000/swagger en tu navegador.

### 🧠 Contribuciones

Este proyecto está en desarrollo. Se aceptan PRs, sugerencias y mejoras.
