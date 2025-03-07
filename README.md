# LicenseToGo

LicenseToGo is a simple, self-hosted license management system for software applications. It allows you to create, validate, and manage license keys through a web interface or API endpoints.

## Overview

This application consists of two main components:

- **Backend API**: Built with PocketBase and Go, handles license generation, validation, and management
- **Web Interface**: Provides a user-friendly interface for license management built with SvelteKit

Data is stored in a local SQLite database that's automatically created in a mounted volume. The application uses the KeysToGo package for the core license management functionality.

## Features

- Create and manage software licenses
- Generate API keys for authenticated access to management endpoints
- Validate licenses via API endpoints
- Self-hosted and containerized solution
- Persistent storage using SQLite
- Simple, containerized deployment

## Quick Start

### Prerequisites

- Docker
- Docker Compose (optional, but recommended)

### Default Credentials and Security Notice

**Important:** LicenseToGo is currently in early development and **not recommended for production use**.

Default admin credentials:

- Email: `admin@admin.com`
- Password: `Pass123!`

You can override these defaults by setting environment variables:

- `USER_EMAIL` - Custom admin email
- `USER_PASSWORD` - Custom admin password

⚠️ **Security Notice:** There is currently no way to change the password after initial setup. This feature is planned for future releases.

### Installation

#### Using Docker

```bash
# Pull the images
docker pull karurosuxx/licensetogo-api:latest
docker pull karurosuxx/licensetogo-web:latest

# Create a network for the containers
docker network create licensetogo-network

# Run the API container
docker run -d \
  --name licensetogo-api \
  -p 8090:8090 \
  -v ${HOME}/ltg/pb_data:/app/pb_data \
  -e USER_EMAIL=your-email@example.com \
  -e USER_PASSWORD=your-secure-password \
  --network licensetogo-network \
  karurosuxx/licensetogo-api:latest

# Run the Web container
docker run -d \
  --name licensetogo-web \
  -p 80:3000 \
  -e PUBLIC_API_URL=http://licensetogo-api:8090 \
  -e ORIGIN=http://localhost:80 \
  --network licensetogo-network \
  karurosuxx/licensetogo-web:latest
```

#### Using Docker Compose

Create a `docker-compose.yml` file:

```yaml
version: "3"
services:
  licensetogo-api:
    image: karurosuxx/licensetogo-api:latest
    container_name: licensetogo-api
    volumes:
      - ${HOME}/ltg/pb_data:/app/pb_data
    ports:
      - "8090:8090"
    networks:
      - licensetogo-network
    # environment:
    # Optional: Custom admin credentials,
    # Once container is ran for the first time, the admin user will be created
    # - USER_EMAIL=your-email@example.com
    # - USER_PASSWORD=your-secure-password
  licensetogo-web:
    image: karurosuxx/licensetogo-web:latest
    container_name: licensetogo-web
    environment:
      - PUBLIC_API_URL=http://licensetogo-api:8090
      - ORIGIN=http://localhost:80
    ports:
      - "80:3000"
    networks:
      - licensetogo-network
networks:
  licensetogo-network:
    name: licensetogo-network
    driver: bridge
```

Then run:

```bash
docker-compose up -d
```

For development, you can use the same Docker Compose file but build the images locally:

```yaml
version: "3"
services:
  licensetogo-api:
    build:
      context: .
      dockerfile: Dockerfile
      target: api
    container_name: licensetogo-api
    volumes:
      - ${HOME}/ltg/pb_data:/app/pb_data
    ports:
      - "8090:8090"
    networks:
      - licensetogo-network
  licensetogo-web:
    build:
      context: .
      dockerfile: Dockerfile
      target: web
    container_name: licensetogo-web
    environment:
      - PUBLIC_API_URL=http://licensetogo-api:8090
      - ORIGIN=http://localhost:80
    ports:
      - "80:3000"
    networks:
      - licensetogo-network
networks:
  licensetogo-network:
    name: licensetogo-network
    driver: bridge
```

Then run:

```bash
docker-compose up -d --build
```

### Usage

1. Access the web interface at `http://localhost:80`
2. Create API keys for application authentication
3. Generate and manage licenses through the UI
4. Integrate your application with the LicenseToGo API

## API Endpoints

LicenseToGo provides the following REST API endpoints for license management:

### License Management

- `GET /api/licenses` - List all licenses

  - Query parameters:
    - `limit` - Maximum number of licenses to return (default: 10)
    - `offset` - Number of licenses to skip (default: 0)
  - Requires API key or user authentication

- `POST /api/licenses` - Create new license

  - Request body:

    ```json
    {
      "name": "License Name",
      "expires": "2025-12-31", // Optional, format: YYYY-MM-DD
      "permissions": [
        // Optional
        "user:read",
        "user:write"
      ],
      "metadata": {
        // Optional
        "customField1": "value1",
        "customField2": "value2"
      }
    }
    ```

  - Requires API key or user authentication

- `PUT /api/licenses/{id}` - Update license

  - Request body:

    ```json
    {
      "name": "Updated License Name", // Optional
      "active": true, // Optional
      "expiresAt": "2025-12-31", // Optional, format: YYYY-MM-DD
      "permissions": [
        // Optional
      ],
      "metadata": {
        // Optional
        "customField1": "updatedValue"
      }
    }
    ```

  - Requires API key or user authentication

- `DELETE /api/licenses/{id}` - Delete license

  - Requires API key or user authentication

- `POST /api/licenses/validate` - Validate a license key

  - Request body:

    ```json
    {
      "key": "your-license-key-here",
      "permissions": [
        // Optional
        "can_create_users",
        "has_module1"
      ]
    }
    ```

  - Public endpoint (no authentication required)

The API is built using [PocketBase](https://pocketbase.io/) and uses the [KeysToGo](https://github.com/karurosux/keystogo) package for license management.

## Technical Details

LicenseToGo is built using:

- **Backend**:

  - [PocketBase](https://pocketbase.io/) - an open-source backend as a service
  - [KeysToGo](https://github.com/karurosux/keystogo) - a custom package for license management
  - Go programming language

- **Frontend**:
  - SvelteKit - a framework for building web applications

The application uses a middleware system to ensure endpoints are protected with either API key or user authentication. The license validation endpoint is the only public endpoint, allowing external applications to verify license keys without authentication.

### License Features

- **Resource-based permissions**: Define what resources and actions a license can access
- **Expiration dates**: Set optional expiration dates for licenses
- **Metadata storage**: Attach custom metadata to licenses
- **Validation**: Simple API for validating licenses with permission checks

## Contributing

Contributions are welcome! This is a personal project in early development stages.

## Status

This project is currently in early development and is very basic. Features and documentation will be expanded over time.

### Known Limitations

- No password change functionality after initial setup
- Limited user management features
- Basic UI with core functionality only
- Not recommended for production environments yet

### Planned Features

- Improved user management
- Password change functionality
- Enhanced security features
- UI improvements
- Comprehensive documentation

## Contributing

Contributions are welcome! This is a personal project in early development stages.

## Status

This project is currently in early development and is very basic. Features and documentation will be expanded over time.

## License

```
MIT License

Copyright (c) 2025 Carlos Gonzalez

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```
