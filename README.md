# Fiber REST API Boilerplate
A boilerplate for [Fiber](https://gofiber.io/) REST API.

## What's included
- SQLite3 database
- [GORM](https://gorm.io/)
- JWT authentication
- Middleware for authorization
- [Air](https://github.com/cosmtrek/air) for hot reloading

## Usage
### 1. Clone the repository
```bash
git clone https://github.com/irfanalmsyah/fiber-boilerplate.git
```

### 2. Change directory
```bash
cd fiber-boilerplate
```

### 3. Install dependencies
```bash
go mod download
```

### 4. Create `.env` file
```bash
cp .env.example .env
```

### 5. Run the application
```bash
air
```

Server will be running on `localhost:3000`.

## TODO
- Tests
- Docker
- Endpoints documentation
- Forgot password
