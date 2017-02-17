package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Bergamo/widgets-spa-go-api/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// UserController represents the controller for operating on the User resource
	UserController struct {
		session *mgo.Session
	}
)

// NewUserController provides a reference to a UserController with provided mongo session
func NewUserController(session *mgo.Session) *UserController {
	return &UserController{session}
}

// Index retrieves a default response
func (userController UserController) Index(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("Hello, World!"))
}

// GetUsers retrieves users
func (userController UserController) GetUsers(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	results := []models.User{}

	// Fetch usernil
	if err := userController.session.DB("widgets-spa-go-api").C("users").Find(nil).All(&results); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	usersJSON, _ := json.Marshal(results)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", usersJSON)
}

// GetUser retrieves an individual user resource
func (userController UserController) GetUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// Grab id
	vars := mux.Vars(r)
	id := vars["id"]

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub user
	user := models.User{}

	// Fetch user
	if err := userController.session.DB("widgets-spa-go-api").C("users").FindId(oid).One(&user); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	userJSON, _ := json.Marshal(user)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", userJSON)
}

// CreateUser creates a new user resource
func (userController UserController) CreateUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Stub an user to be populated from the body
	user := models.User{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&user)

	// Add an Id
	user.Id = bson.NewObjectId()

	// Write the user to mongo
	userController.session.DB("widgets-spa-go-api").C("users").Insert(user)

	// Marshal provided interface into JSON structure
	userJSON, _ := json.Marshal(user)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", userJSON)
}

// UpdateUser updates user
func (userController UserController) UpdateUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Stub an user to be populated from the body
	user := models.User{}

	// Grab id
	vars := mux.Vars(r)
	id := vars["id"]

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Populate the user data
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Write the user to mongo
	if err := userController.session.DB("widgets-spa-go-api").C("users").UpdateId(oid, &user); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	userJSON, _ := json.Marshal(user)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", userJSON)
}

// RemoveUser removes an existing user resource
func (userController UserController) RemoveUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// Grab id
	vars := mux.Vars(r)
	id := vars["id"]

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err := userController.session.DB("widgets-spa-go-api").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	// Write status
	w.WriteHeader(200)
}
