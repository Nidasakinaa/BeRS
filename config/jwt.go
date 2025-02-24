package config

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt/v4"
	"github.com/Nidasakinaa/BeRS/model"
	 "golang.org/x/crypto/bcrypt"
)

func GenerateJWT(admin model.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["admin_id"] = admin.ID.Hex()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := os.Getenv("JWT_SECRET")	
	return token.SignedString([]byte(secretKey))
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}