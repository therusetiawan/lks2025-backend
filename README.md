# Golang Simple CRUD App with AWS S3, Redis, and Docker

This is a simple CRUD application built with Golang, using:
- **MariaDB (AWS RDS)** as the database
- **Redis (AWS ElastiCache)** for caching
- **AWS S3** for storing uploaded images
- **Docker & Docker Compose** for containerized deployment

## Features
- CRUD operations (Create, Read, Update, Delete)
- Uses **GORM** ORM for database interactions
- Uploads images to **AWS S3**
- Caches data in **AWS ElastiCache (Redis)**
- Deployable with **Docker Compose**

## Requirements
- Golang 1.20+
- Docker & Docker Compose
- AWS RDS (MariaDB)
- AWS ElastiCache (Redis)
- AWS S3 Bucket & IAM Credentials

---



---

## 1️⃣ Project Setup

### 1.1 Clone the Repository
```sh
git clone https://github.com/therusetiawan/lks2025cc-backend.git
cd golang-crud-app
```

### 1.2 Update Environment Variables
Modify `docker-compose.yml`:
```yaml
services:
  app:
    build: .
    container_name: golang_crud_app
    ports:
      - "8080:8080"
    environment:
      - DB_DSN=user:password@tcp(aws-rds-endpoint:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local
      - REDIS_ADDR=aws-redis-endpoint:6379
      - AWS_REGION=us-east-1
      - AWS_BUCKET_NAME=my-golang-app-bucket
      - AWS_ACCESS_KEY=your-access-key
      - AWS_SECRET_KEY=your-secret-key
```

### 1.3 Build and Run the App
```sh
docker-compose up --build -d
```

---

## 2️⃣ API Endpoints

| Method | Endpoint       | Description |
|--------|---------------|-------------|
| POST   | `/users`      | Create a new user |
| GET    | `/users`      | Get all users |
| GET    | `/users/{id}` | Get a user by ID |
| PUT    | `/users/{id}` | Update a user |
| DELETE | `/users/{id}` | Delete a user |

### 2.1 Upload File to S3
When adding a user, the **photo** field supports file uploads:
- The file is uploaded to **AWS S3**.
- The S3 URL is stored in the database.

### 2.2 Stop the Application
```sh
docker-compose down
```

---

## 3️⃣ License
MIT License

## 4️⃣ Author
Your Name - [GitHub Profile](https://github.com/therusetiawan)
