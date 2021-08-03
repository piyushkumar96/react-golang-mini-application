package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pascaldekloe/jwt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"react-golang-mini-application/movies-app-backend/models"
	"time"
)

var validUser = models.User{
	ID:       10,
	Email:    "piyush25032@gmail.com",
	Password: "$2a$12$lkQf3tYuPpNuQvQqwcQrZOermijz0trW8nhL8FJPi0VXTxXIM9n3q",
}

type Credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

func (app *application) signIn(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"))
		return
	}
	hashedPassword := validUser.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password))
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"))
		return
	}

	var claims jwt.Claims
	claims.Subject = fmt.Sprint(validUser.ID)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))
	claims.Issuer = "github.com/piyushkumar96"
	claims.Audiences = []string{"github.com/piyushkumar96"}

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(app.config.jwt.secret))
	if err != nil {
		app.errorJSON(w, errors.New("error signing"))
		return
	}

	err = app.writeJSON(w, http.StatusOK, string(jwtBytes), "response")
	if err != nil {
		app.errorJSON(w, errors.New("error signing"))
		return
	}
}
