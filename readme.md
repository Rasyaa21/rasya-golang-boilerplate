# Rasya Golang Boilerplate

A simple **Golang** boilerplate using **Gin** and **MySQL**, designed for clean and scalable API development.

---

## 🚀 Features
- REST API with **Gin**
- Database integration using **GORM** and **MySQL**
- **Environment variables** with `.env` support using `godotenv`
- Organized **folder structure**

---

## 🛠️ Installation and Setup

### 1️⃣ **Clone the Repository**
```sh
git clone https://github.com/your-username/rasya-golang-boilerplate.git
cd <project-name>

---

Ensure **Go is installed** (check with `go version`). If not installed, [download Go](https://go.dev/dl/) first.

Then, install dependencies:

```sh
go mod init <project-name>
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/mysql
go get github.com/joho/godotenv