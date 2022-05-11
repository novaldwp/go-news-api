package helper

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var JWT_SIGNATURE_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))    // set signature key
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256                // using HS256 algorithms for hash
var LOGIN_EXPIRATION_DURATION = time.Now().Add(24 * time.Hour) // set token expire 24hour

type JWTClaim struct { // struct for claim
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(email, role string) (tokenString string, err error) {
	expirationTime := LOGIN_EXPIRATION_DURATION
	claims := &JWTClaim{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)
	tokenString, err = token.SignedString(JWT_SIGNATURE_KEY)

	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT_SIGNATURE_KEY), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)

	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	return
}
