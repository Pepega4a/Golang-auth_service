package services

import (
	"encoding/base64"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GenerateRefreshToken(ip string) (string, string, error) {
	rawToken := fmt.Sprintf("%d:%s", time.Now().UnixNano(), ip)
	encodedToken := base64.StdEncoding.EncodeToString([]byte(rawToken))
	hash, err := bcrypt.GenerateFromPassword([]byte(encodedToken), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}
	return encodedToken, string(hash), nil
}

func VerifyRefreshToken(hash, token string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(token))
}
