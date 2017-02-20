package controllers

import (
	"net/http"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"fmt"
	request "github.com/dgrijalva/jwt-go/request"
)


type (
	// TokenController represents the controller for operating on the Token resource
	TokenController struct {
		mySigningKey []byte
	}
)

/* Set up a global string for our secret */
var mySigningKey = []byte("secret")

// NewTokenController provides a reference to a TokenController
func NewTokenController() *TokenController {
	return &TokenController{mySigningKey}
}

// ValidateTokenHandler validate token from request
func (tokenController TokenController) ValidateTokenHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return tokenController.mySigningKey, nil
		})

	if err == nil {
		if token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
	}
}

// GetTokenHandler retrieves a token
func (tokenController TokenController) GetTokenHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["admin"] = true
	claims["name"] = "Ado Kukic"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(tokenController.mySigningKey)

	/* Finally, write the token to the browser window */
	w.Write([]byte(tokenString))
}
