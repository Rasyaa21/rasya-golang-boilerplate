# Project Name

A simple web application using **Gin**, **GORM**, and **MySQL** in Go.

## Prerequisites

Make sure you have **Go** installed. You can download it from [golang.org](https://golang.org/dl/).

## Installation

Follow these steps to set up the project:

### 1. Initialize Go Module
```sh
 go mod init <project-name>
```
Replace `<project-name>` with your actual project name.

### 2. Install Dependencies
```sh
 go get github.com/gin-gonic/gin
 go get gorm.io/gorm
 go get gorm.io/driver/mysql
 go get github.com/joho/godotenv
```

### 3. Set Up Environment Variables
Create a `.env` file in the project root and add your MySQL database credentials:
```env
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=yourdatabase
DB_HOST=localhost
DB_PORT=3306
```

### 4. Create Main Application File
Create a `main.go` file and set up a basic Gin server:
```go
package main

import (
    "fmt"
    "log"
    "os"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
)

type User struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string `gorm:"size:100"`
    Email string `gorm:"unique"`
}

var DB *gorm.DB

func initDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
    
    var errDB error
    DB, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if errDB != nil {
        log.Fatal("Failed to connect to database")
    }
    DB.AutoMigrate(&User{})
}

func main() {
    initDB()
    r := gin.Default()
    
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello, World!"})
    })
    
    r.Run(":8080")
}
```

### 5. Run the Application
```sh
 go run main.go
```

## API Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST    | `/api/auth/register`      | Register User |
| POST    | `/api/auth/login`      | Login User |
| GET    | `/api/user/`      | Get Current User Detail |

## License
This project is licensed under the MIT License.
