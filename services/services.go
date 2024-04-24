package services

import (
	"context"
	"fmt"
	"github.com/Venukishore-R/microservice1_auth/models"
	"github.com/go-kit/log"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"net/http"
	"os"
	"time"
)

type LoggerService struct {
	logger log.Logger
}

type Service interface {
	Register(ctx context.Context, name, email, phone, password string) (int64, string, error)
	Login(ctx context.Context, email, password string) (int64, string, string, error)
}

func NewLoggerService(logger log.Logger) *LoggerService {
	return &LoggerService{
		logger: logger,
	}
}

func ConnectDb() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (s *LoggerService) Register(ctx context.Context, name, email, phone, password string) (int64, string, error) {
	s.logger.Log("received:", "name", name, "email", email, "phone", phone, "password", password)

	db, err := ConnectDb()
	if err != nil {
		s.logger.Log("error while connecting to db:", err)
		return http.StatusInternalServerError, "unable to connect to db", err
	}

	user := &models.User{
		Name:     name,
		Phone:    phone,
		Email:    email,
		Password: password,
	}

	err = db.Model(&models.User{}).Create(user).Error
	if err != nil {
		s.logger.Log("unable to register user:", err)
		return http.StatusInternalServerError, "unable to register user", err
	}

	return http.StatusOK, "user registered", nil
}

func (s *LoggerService) Login(ctx context.Context, email, password string) (int64, string, string, error) {
	s.logger.Log("received: ", "email", email, "password", password)

	db, err := ConnectDb()
	if err != nil {
		s.logger.Log("unable to connect to db", err)
		return http.StatusInternalServerError, "", "", err
	}

	var user *models.User

	err = db.Model(&models.User{}).Where("email=?", email).First(&user).Error
	if err != nil {
		s.logger.Log("no user", err)
		return http.StatusUnauthorized, "", "", err
	}

	if user.Password != password {
		s.logger.Log("password mismatch")
		return http.StatusUnauthorized, "", "", fmt.Errorf("password mismatch")
	}

	claims := &models.UserClaims{
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	refreshClaims := &models.UserClaims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken.Header["kid"] = "access_token"

	newAccessToken, err := accessToken.SignedString(models.JwtUserKey)
	if err != nil {
		s.logger.Log("unable to create access token", err)
		return http.StatusInternalServerError, "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken.Header["kid"] = "refresh_token"

	newRefreshToken, err := refreshToken.SignedString(models.JwtUserKey)
	if err != nil {
		s.logger.Log("unable to create refresh token", err)
		return http.StatusInternalServerError, "", "", err
	}

	return http.StatusOK, newAccessToken, newRefreshToken, nil
}
