package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"hack-change-api/models/auxiliary"
	u "hack-change-api/muxutil"
	"net/http"
	"os"
	"strings"
)

// JwtAuth is validating JWT-token attached in Authorization header.
var JwtAuth = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// list of endpoints that does NOT require auth
		notAuth := []string{"/api/auth/login", "/api/auth/register"}
		requestPath := r.URL.Path // current request path

		// check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization") // grab the token from the header

		if tokenHeader == "" { // token is missing, returns with error code 401 Unauthorized
			u.HandleUnauthorized(w, errors.New("token is not provided"))
			return
		}

		// the token normally comes in format `Bearer {token-body}`
		// we just check if the retrieved token matched this requirement
		split := strings.Split(tokenHeader, " ")
		if len(split) != 2 {
			u.HandleUnauthorized(w, errors.New("invalid/malformed auth token: " + tokenHeader))
			return
		}

		tokenPart := split[1] // grab the token part, what we are truly interested in
		tk := &auxiliary.JWT{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil { // malformed token, returns with http code 403 as usual
			u.HandleUnauthorized(w, errors.New("malformed authentication token: " + err.Error()))
			return
		}

		if !token.Valid { // token is invalid, maybe not signed on this server
			u.HandleUnauthorized(w, errors.New("token is not valid"))
			return
		}

		go updateLastVisit(tk.UserID)

		// everything went well, proceed with the request
		// and set the caller to the user retrieved from the parsed token
		v := u.Values{M: map[string]string {
			"user": fmt.Sprint(tk.UserID),
		}}

		ctx := context.WithValue(r.Context(), "context", v)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r) // proceed in the middleware chain!
	})
}

func updateLastVisit(userID uint) {
	log.Debug("User ID: ", userID) // useful for monitoring
	//err := db.GetDB().Model(&entities.User{}).
	//	Where("id = ?", userID).
	//	Update("last_visit", time.Now()).Error
	//
	//if err != nil {
	//	log.Error(err)
	//}
}
