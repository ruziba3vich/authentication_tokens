package jwttokens

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func ExtractUserIDFromToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, nil
	} else {
		return 0, fmt.Errorf("invalid token")
	}
}
