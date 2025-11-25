# 30Budget - Personal Budget Management App

A modern, AI-powered personal budget management application built with React (frontend) and Go (backend).

## Project Structure

```
30budget/
├── frontend/           # React + Vite + TypeScript SPA
├── backend/            # Go REST API with PostgreSQL
├── docker-compose.yml  # Docker Compose for full-stack setup
└── .env.example        # Environment variables template
```

## Prerequisites

- Docker & Docker Compose (for containerized setup)
- Node.js 20+ (for local frontend development)
- Go 1.24+ (for local backend development)
- PostgreSQL 16+ (for local backend development)

## Quick Start with Docker Compose

### 1. Clone or setup environment

```bash
cd /path/to/30budget
cp .env.example .env
```

Edit `.env` with your values (at minimum, change `JWT_SECRET` for production).

### 2. Build and run all services

```bash
docker-compose up --build
```

This will start:
- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **PostgreSQL**: localhost:5432
- **Redis**: localhost:6379

### 3. Access the application

Open your browser to [http://localhost:3000](http://localhost:3000)

### 4. Stop services

```bash
docker-compose down
```

To also remove database volume:

```bash
docker-compose down -v
```

---

## Local Development (without Docker)

### Backend Setup

```bash
cd backend

# Set environment variables
export DATABASE_URL="postgres://postgres:password@localhost:5432/budget_db?sslmode=disable"
export JWT_SECRET="dev-secret-key"
export CORS_ORIGINS="http://localhost:3000"
export ENVIRONMENT="development"

# Start PostgreSQL (if not using Docker)
# docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=password postgres:16-alpine

# Run migrations
go run . # or build and run binary

# Start the server (assumes migrate and seedData are called)
```

### Frontend Setup

```bash
cd frontend

# Install dependencies
npm install

# Set backend API URL
export VITE_API_URL=http://localhost:8080

# Start dev server
npm run dev

# Open http://localhost:5173 (Vite default port)
```

---

## API Endpoints

### Authentication

- `POST /api/v1/auth/signup` - Create a new user
- `POST /api/v1/auth/login` - User login (returns access token, sets refresh cookie)
- `POST /api/v1/auth/refresh` - Refresh access token
- `POST /api/v1/auth/logout` - User logout

### User

- `GET /api/v1/me` - Get current user info
- `PATCH /api/v1/me` - Update user profile

### Budget & Settings

- `GET /api/v1/settings` - Get user settings (budget, currency, etc.)
- `PATCH /api/v1/settings` - Update user settings

### Categories

- `GET /api/v1/categories` - List all categories
- `POST /api/v1/categories` - Create new category
- `PATCH /api/v1/categories/{id}` - Update category
- `DELETE /api/v1/categories/{id}` - Delete category

### Transactions

- `GET /api/v1/transactions?from_date=&to_date=&limit=&offset=` - List transactions
- `POST /api/v1/transactions` - Create transaction
- `PATCH /api/v1/transactions/{id}` - Update transaction
- `DELETE /api/v1/transactions/{id}` - Delete transaction

### Notifications

- `GET /api/v1/notifications?limit=&offset=` - Get notifications
- `POST /api/v1/notifications/mark-read` - Mark notifications as read
- `POST /api/v1/notifications/mark-all-read` - Mark all as read

---

## Database Schema

Core tables:
- `users` - User accounts with settings
- `categories` - Income/expense categories
- `transactions` - Financial transactions
- `notifications` - User notifications
- `templates` - Budget templates
- `template_categories` - Categories within templates
- `refresh_tokens` - Refresh token storage for auth
- `user_settings` - User preferences

---

## Frontend Architecture

### Contexts

- **AuthContext** (`context/AuthContext.tsx`) - Authentication state and methods
- **BudgetContext** (`context/BudgetContext.tsx`) - Budget data and CRUD operations
- **ThemeContext** (`context/ThemeContext.tsx`) - Dark/light theme management

### Key Components

- **Dashboard** - Overview of budget status
- **Budgeting** - Set and track budget limits
- **Transactions** - Add/edit/view transactions
- **Categories** - Manage transaction categories
- **Analytics** - Budget charts and insights
- **Settings** - User preferences
- **Templates** - Budget templates

---

## Backend Architecture

### Layers

```
api/handlers/       - HTTP request handlers
api/dto/           - Request/Response data transfer objects
api/middleware/    - Auth, CORS, logging, rate limiting
service/           - Business logic
storage/           - Database access layer
db/queries/        - SQL queries (sqlc-managed)
db/sqlc/           - Generated types from sqlc
auth/              - JWT and refresh token services
config/            - Configuration management
utils/             - Helpers and utilities
```

### Tech Stack

- **Framework**: Chi v5 (lightweight HTTP router)
- **Database**: PostgreSQL with pgx driver
- **Auth**: JWT + refresh tokens in httpOnly cookies
- **Password**: bcrypt hashing
- **Logging**: Zap structured logging
- **Cache**: Redis
- **Email**: Brevo API

---

## Environment Variables

### Required

```
DATABASE_URL          # PostgreSQL connection string
JWT_SECRET           # Secret key for signing tokens
CORS_ORIGINS         # Comma-separated allowed origins
```

### Optional

```
ENVIRONMENT          # development | production (default: development)
PORT                 # Backend port (default: 8080)
LOG_LEVEL            # debug | info | warn | error (default: info)
BREVO_API_KEY        # Email service API key
```

See `.env.example` for all options.

---

## Authentication Flow

### Signup/Login

1. User submits email + password → Backend
2. Backend validates, hashes password (bcrypt), creates user record
3. Backend generates:
   - **Access Token**: Short-lived (15 min), stored in memory by SPA, sent in `Authorization: Bearer <token>` header
   - **Refresh Token**: Long-lived (7 days), stored in httpOnly cookie, sent automatically with requests
4. Backend returns `{ access_token, token_type, expires_in, user }`
5. SPA stores access token in memory, receives refresh cookie via Set-Cookie

### Protected Requests

- SPA sends `Authorization: Bearer <access_token>` with each API request
- Backend middleware validates token
- If expired, frontend calls `/auth/refresh` endpoint
- Backend reads httpOnly cookie, issues new access token
- SPA updates memory and retries original request

### Logout

- User clicks logout → Frontend calls `POST /api/v1/auth/logout`
- Backend revokes refresh token, clears cookie
- Frontend clears memory state

---

## Security Considerations

- **HTTPS Required** in production (use reverse proxy like Nginx)
- **CSRF Protection** via SameSite cookies + CORS validation
- **Rate Limiting** on auth endpoints (built-in via middleware)
- **Input Validation** on all API endpoints
- **SQL Injection Prevention** via parameterized queries
- **Sensitive Data**: Never log passwords, tokens, or PII

---

## Development Workflow

### Making Backend Changes

1. Edit service/handler/storage files
2. Add/update SQL queries in `internal/db/queries/*.sql`
3. Run `sqlc generate` to create types (if using sqlc)
4. Test via `go test ./...` or make API requests
5. Rebuild Docker image: `docker-compose build backend`
6. Restart: `docker-compose up backend`

### Making Frontend Changes

1. Edit React components/contexts
2. Frontend dev server hot-reloads automatically
3. Test in browser at http://localhost:3000
4. Build: `npm run build` (in `frontend/` dir)

---

## Troubleshooting

### Database Connection Failed

```bash
# Check if PostgreSQL is running
docker-compose ps db

# View logs
docker-compose logs db

# Restart database
docker-compose restart db
```

### Backend won't start

```bash
# Check logs
docker-compose logs backend

# Verify migrations ran
docker-compose exec backend ./main --migrate

# Reset DB (caution!)
docker-compose down -v && docker-compose up --build
```

### Frontend not connecting to backend

- Check `VITE_API_URL` matches backend address
- Verify CORS origins in backend config
- Check browser console for fetch errors
- Ensure backend is healthy: `docker-compose ps backend`

---

## Performance Tips

- Use Redis for caching frequently-accessed data
- Paginate large result sets (transactions, notifications)
- Index frequently-queried columns (user_id, date)
- Use connection pooling (configured in DB connection)

---

## Next Steps

- [ ] Add email notifications (Brevo integration)
- [ ] Implement budget alerts (server-side scheduled jobs)
- [ ] Add data export (CSV, PDF)
- [ ] Implement recurring transactions
- [ ] Add mobile app support
- [ ] Deploy to cloud (AWS, GCP, DigitalOcean)

---

## Support & Contributions

For issues or questions, please open a GitHub issue.

---

## License

MIT
