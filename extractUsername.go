package jwttokens

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func ExtractUsernameFromToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Username, nil
	} else {
		return "", fmt.Errorf("invalid token")
	}
}
