package handler

import (
	"time"

	"github.com/SawitProRecruitment/UserService/repository"
	
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const signingSecret = "thisisaverylongbutsecuretokenstring"

type claims struct {
	UserUUID uuid.UUID `json:"user_uuid"`
	jwt.StandardClaims
}

func NewBearerToken(user *repository.Users) (string, time.Time) {
	expiry := time.Now().Add(time.Hour * 24).Unix()
	claims := &claims{
		UserUUID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiry,
		},
	}

	// if you want to use RS256 use this function and your private key
	// privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(myPrivateKey)
	// token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SigningString(privateKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenString, err := token.SignedString([]byte(signingSecret)); err == nil {
		return tokenString, time.Unix(expiry, 0)
	}
	return "", time.Unix(expiry, 0)
}
