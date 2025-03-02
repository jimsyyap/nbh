# Tennis Club Blog

Welcome to the Tennis Club Blog project! This is a web application designed to manage and display blog content for a tennis club, featuring user authentication, content management, and an admin panel. The backend is built with Go, and the frontend is developed using React.js, with a PostgreSQL database for persistence.

## Table of Contents
- [Project Overview](#project-overview)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Project Structure](#project-structure)
- [Development Phases](#development-phases)
- [Running the Application](#running-the-application)
- [Contributing](#contributing)
- [License](#license)

## Project Overview
The Tennis Club Blog is a full-stack web application aimed at providing a platform for tennis club members to read, create, and manage blog posts. It includes a public-facing blog, user authentication, and an admin panel for content management.

- **Backend**: Built with Go, featuring RESTful APIs, PostgreSQL integration, and JWT-based authentication.
- **Frontend**: Developed with React.js, leveraging Redux/Context API for state management and responsive design.
- **Database**: PostgreSQL for storing users, posts, comments, and media.

## Features
- User authentication and authorization
- Blog post creation, editing, and deletion
- Commenting system
- Media upload and management
- Responsive UI for desktop and mobile
- Admin panel for content moderation

## Prerequisites
- Go (version 1.21 or later)
- Node.js (version 18 or later)
- PostgreSQL (version 15 or later)
- Git
- Docker (optional, for containerized deployment)

## Installation

### Backend
1. Clone the repository:
   ```bash
   git clone https://github.com/your-org/tennis-club-blog.git
   cd tennis-club-blog/tennis-club-api
   ```
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Copy the example environment file and configure it:
   ```bash
   cp .env.example .env
   ```
   Edit `.env` with your database credentials and JWT secret.
4. Set up the PostgreSQL database and run migrations:
   ```bash
   psql -U your_user -d your_db -f migrations/0001_create_tables.sql
   ```
5. Build and run:
   ```bash
   go build -o tennis-club-api cmd/api/main.go
   ./tennis-club-api
   ```

### Frontend
1. Navigate to the frontend directory:
   ```bash
   cd tennis-club-blog/tennis-club-client
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Copy the example environment file and configure it:
   ```bash
   cp .env.example .env
   ```
   Update `.env` with the backend API URL.
4. Start the development server:
   ```bash
   npm start
   ```

## Project Structure

### Backend (Go)
```
tennis-club-api/
├── cmd/api/main.go                # Application entry point
├── internal/                      # Private application code
│   ├── config/                    # Configuration handling
│   ├── domain/                    # Domain models (user, post, comment, media)
│   ├── repository/                # Database access layer
│   ├── service/                   # Business logic
│   ├── delivery/http/             # HTTP handlers and routing
│   └── utils/                     # Utility functions
├── migrations/                    # Database migrations
├── pkg/                           # Public packages (logger, validator, jwt)
├── scripts/                       # Build/deployment scripts
├── test/                          # Test utilities
├── .env.example                   # Environment variables template
├── Dockerfile                     # Docker configuration
├── go.mod                         # Go module file
└── go.sum                         # Go dependencies
```

### Frontend (React.js)
```
tennis-club-client/
├── public/                        # Static files
├── src/
│   ├── assets/                    # Images, fonts, etc.
│   ├── components/                # Reusable UI components
│   ├── pages/                     # Page components (Home, BlogPost, Admin)
│   ├── services/                  # API client services
│   ├── store/                     # State management (Redux/Context)
│   ├── hooks/                     # Custom React hooks
│   ├── utils/                     # Utility functions
│   ├── context/                   # React Context
│   ├── routes/                    # Application routing
│   ├── App.js                     # Main app component
│   └── index.js                   # Entry point
├── .env.example                   # Environment variables template
├── package.json                   # NPM configuration
├── Dockerfile                     # Docker configuration
└── README.md                      # Frontend-specific docs
```

## Development Phases
1. **Phase 1: Project Setup and Planning** (2-4 weeks)
   - Requirements analysis, technical setup, and database design.
2. **Phase 2: Backend Development** (4-6 weeks)
   - API development, database implementation, and testing.
3. **Phase 3: Frontend Development** (4-6 weeks)
   - UI implementation, state management, and testing.
4. **Phase 4: Integration and Deployment** (2-3 weeks)
   - End-to-end testing, production deployment, and documentation.
5. **Phase 5: Maintenance and Iteration** (Ongoing)
   - Monitoring, bug fixes, and feature updates based on feedback.

## Running the Application

### Locally
- Start the backend: `go run cmd/api/main.go`
- Start the frontend: `npm start` (from `tennis-club-client/`)
- Access the app at `http://localhost:3000`

### With Docker
1. Build and run with Docker Compose:
   ```bash
   docker-compose up --build
   ```
2. Access the app at `http://localhost:3000`

## Contributing
1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -m "Add your feature"`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a Pull Request.

Please follow the coding standards and include tests for new features.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
