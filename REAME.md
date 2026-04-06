# nexo

nexo is a local-first self-hosted app for habits, reading backlog, notes, and daily overview.

## Status
This project is in early development.

## Planned MVP
- Authentication
- Dashboard
- Habit tracking
- Reading library
- Notes
- Responsive web UI

## Tech Stack
- Go
- PostgreSQL
- Docker Compose
- Server-rendered HTML

## Getting Started

### Requirements
- Docker
- Docker Compose

### Run locally
```bash
cp .env.example .env
docker compose up --build
```

The app will be available at:
- http://localhost:8080

## Project Structure
- `cmd/api` - application entrypoint
- `internal` - core application code
- `web` - templates and static assets
- `migrations` - database migrations

## License
This project is source-available under the PolyForm Noncommercial 1.0.0 license.
Commercial use, resale, and paid hosting are not permitted.