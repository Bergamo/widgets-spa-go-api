package authentication

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

// RequireTokenAuthentication asd
func RequireTokenAuthentication(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

}

/* Set up a global string for our secret */
var mySigningKey = []byte("secret")

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
