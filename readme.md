# TaskNest

**TaskNest** is a lightweight task management web application built with Go and Fiber. It allows users to register, log in, and manage their personal todos with a clean HTML interface and optional API access.

---

## ✨ Features

- 🔐 User authentication (register, login, logout)
- ✅ Create, read, update, and delete todos
- 🖼️ Responsive HTML views for todos, login, and registration
- 🧾 RESTful API endpoints for todos
- 🔑 JWT-based session authentication
- 🐘 PostgreSQL database integration

---

## 📁 Project Structure

```text
.env                # Environment variables  
go.mod              # Go module dependencies  
main.go             # Application entry point  
db/                 # Database initialization  
handlers/           # Request handlers  
middlewares/        # Authentication middleware  
models/             # Data models  
routes/             # Route definitions  
utils/              # Utility functions  
views/              # HTML templates  
```

---

## ⚙️ Prerequisites

- Go 1.20+
- PostgreSQL

---

## 🚀 Getting Started

### 🔧 Local Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/aryanpnd/Tasknest.git
   cd Tasknest
   ```

2. Set up your `.env` file with your PostgreSQL credentials:

   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=root
   DB_NAME=tasknest_db
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Run the application:

   ```bash
   go run main.go
   ```

5. Open your browser and navigate to [http://localhost:8080](http://localhost:8080)

---

### 🐳 Optional: Docker Setup

#### Standalone Container

1. Build the Docker image:

   ```bash
   docker build -t tasknest .
   ```

2. Run the container:

   ```bash
   docker run -p 8080:8080 --env-file .env tasknest
   ```

#### Using Docker Compose (with PostgreSQL)

1. Ensure `docker-compose.yml` is present (already included in the repo).

2. Start the services:

   ```bash
   docker-compose up --build
   ```

3. Visit [http://localhost:8080](http://localhost:8080)

---

## 🔗 API & Page Routes

### 👤 User Routes

| Method | Endpoint    | Description             |
|--------|-------------|-------------------------|
| GET    | `/register` | Show registration page  |
| POST   | `/register` | Register a new user     |
| GET    | `/login`    | Show login page         |
| POST   | `/login`    | Log in a user           |

### ✅ Todo Routes

| Method | Endpoint       | Description              |
|--------|----------------|--------------------------|
| GET    | `/todos/html`  | Render todos page        |
| GET    | `/todos`       | Fetch all todos (API)    |
| POST   | `/todos`       | Create a new todo        |
| PUT    | `/todos/:id`   | Update an existing todo  |
| DELETE | `/todos/:id`   | Delete a todo            |

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).