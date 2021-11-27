package auth

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"hack-change-api/db"
	"hack-change-api/hashutil"
	"hack-change-api/models/auxiliary"
	"hack-change-api/models/entities"
	u "hack-change-api/muxutil"
	"net/http"
	"os"
	"time"
)

func login(email, password string) map[string]interface{} {
	account := &entities.User{}
	err := db.GetDB().Table("users").Where("email = ?", email).First(account).Error

	if err != nil {
		log.Warn(err)
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "User not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	if !hashutil.CheckPasswordHash(password, account.Password) { // Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again")
	}

	// worked! logged in
	db.GetDB().Model(&account).Update("LastLogin", time.Now())

	// create JWT token
	tk := &auxiliary.JWT{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))

	resp := u.Message(true, "Logged In")
	resp["token"] = tokenString
	resp["user_id"] = account.ID

	return resp
}

var Login = func(w http.ResponseWriter, r *http.Request) {
	account := &auxiliary.LoginCredentials{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	resp := login(account.Email, account.Password)

	if resp["token"] == nil {
		u.HandleBadRequest(w, errors.New("wrong credentials"))
		return
	}

	u.Respond(w, resp)
}

var Register = func(w http.ResponseWriter, r *http.Request) {
	account := &auxiliary.RegisterCredentials{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	passwordHash, _ := hashutil.HashPassword(account.Password)
	user := entities.User{
		Email:      account.Email,
		Password:   passwordHash,
		Name:       account.Name,
		Surname:    account.Surname,
		IsVerified: false,
	}

	err = db.GetDB().Create(&user).Error
	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(user)
		u.RespondJSON(w, res)
	}
}
