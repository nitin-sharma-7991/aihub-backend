# Environment Configuration

Version: v0.1.0

Status: Stable

---

## Create .env

Create a `.env` file in the project root.

Example

```env
APP_NAME=AIHub
APP_ENV=development
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=aihub
DB_SSLMODE=disable
DB_TIMEZONE=Asia/Kolkata
```

---

## Security

Never commit the `.env` file to version control.