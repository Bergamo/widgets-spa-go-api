package controllers

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

type (
	// TokenController represents the controller for operating on the Token resource
	TokenController struct {
		mySigningKey []byte
	}
)

// NewTokenController provides a reference to a TokenController
func NewTokenController(mySigningKey []byte) *TokenController {
	return &TokenController{mySigningKey}
}

// GetTokenHandler retrieves a token
func (tokenController TokenController) GetTokenHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

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
