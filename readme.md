# 🐳 Dockerized Golang Form + MySQL Project

A simple web form built with Golang that collects user `username` and `email`, stores the data in a MySQL database, and displays all submissions. Everything runs inside Docker containers.

---

## 📦 Tech Stack

- [Golang](https://golang.org/)
- [MySQL](https://www.mysql.com/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## 📁 Project Structure

├── docker-compose.yml
├── init.sql 
├── web 
├── Dockerfile 
├── go.mod 
└── main.go


## 🚀 Getting Started

### 1. Clone the Repository

git clone https://github.com/heinthant/docker-golang-form.git
cd docker-golang-form
2. Run with Docker Compose
docker-compose up --build
This will:

Build the Go app (web service)

Start a MySQL database (db service)

Run DB init script (init.sql) to create the table

🌐 Access the App
Visit: http://localhost:8080

Submit your name and email

View all submitted entries in a styled table

🛢️ Database Info
Container name: docker-golang_form-db-1

Database: userdb

Table: users

User: root

Password: rootpassword

To access MySQL CLI inside the container:


docker exec -it docker-golang_form-db-1 mysql -u root -p
# Enter password: rootpassword
🧹 Stop & Clean Up
docker-compose down
💡 Future Ideas
Add form validation

Search/filter user list

Export to CSV

Deploy to cloud

📝 License
MIT – free for personal and commercial use.
