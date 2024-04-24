package models

import "github.com/golang-jwt/jwt/v4"

type User struct {
	Id       int64  `json:"id" gorm:"primary_key;autoIncrement"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Phone    string `json:"phone" gorm:"unique"`
	Password string `json:"password"`
}

type UserClaims struct {
	Name  string
	Email string
	Phone string
	jwt.StandardClaims
}

var JwtUserKey = []byte("i am the user")
