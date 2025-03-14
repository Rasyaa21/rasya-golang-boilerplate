package models

import (
	"errors"
	"fmt"
	"rasya-golang-boilerplate/config"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Bio       *string   `gorm:"type:text" json:"bio"`
	AvatarURL *string   `gorm:"type:varchar(255)" json:"avatar_url"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func StoreUser(name, email, password string) (*User, error) {
	if config.DB == nil {
		return nil, gorm.ErrInvalidDB
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}
	user := User{Name: name, Email: email, Password: hashedPassword}

	res := config.DB.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	fmt.Println("✅ User successfully stored in database")
	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (user *User) CheckPassword(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func GetUserByID(uid uint) (User, error) {

	var u User

	if err := config.DB.First(&u, uid).Error; err != nil {
		return u, errors.New("user not found")
	}

	u.PrepareGive()

	return u, nil
}

func (u *User) PrepareGive() {
	u.Password = ""
}
