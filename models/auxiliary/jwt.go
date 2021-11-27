package auxiliary

import "github.com/dgrijalva/jwt-go"

// JWT-token claims
type JWT struct {
	UserID  uint
	jwt.StandardClaims
}
