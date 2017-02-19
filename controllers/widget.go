package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/Bergamo/widgets-spa-go-api/models"
	"github.com/Bergamo/widgets-spa-go-api/common"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// WidgetController represents the controller for operating on the Widget resource
	WidgetController struct {
		session *mgo.Session
	}
)

// NewWidgetController provides a reference to a WidgetController with provided mongo session
func NewWidgetController(session *mgo.Session) *WidgetController {
	return &WidgetController{session}
}

// GetWidgets retrieves widgets
func (widgetController WidgetController) GetWidgets(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	results := []models.Widget{}

	// Fetch widgetnil
	if err := widgetController.session.DB(common.AppConfig.Database).C("widgets").Find(nil).All(&results); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	widgetsJSON, _ := json.Marshal(results)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", widgetsJSON)
}

// GetWidget retrieves an individual widget resource
func (widgetController WidgetController) GetWidget(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

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

	// Stub widget
	widget := models.Widget{}

	// Fetch widget
	if err := widgetController.session.DB(common.AppConfig.Database).C("widgets").FindId(oid).One(&widget); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	widgetJSON, _ := json.Marshal(widget)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", widgetJSON)
}

// CreateWidget creates a new widget resource
func (widgetController WidgetController) CreateWidget(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Stub an widget to be populated from the body
	widget := models.Widget{}

	// Populate the widget data
	json.NewDecoder(r.Body).Decode(&widget)

	// Add an Id
	widget.Id = bson.NewObjectId()

	// Write the widget to mongo
	widgetController.session.DB(common.AppConfig.Database).C("widgets").Insert(widget)

	// Marshal provided interface into JSON structure
	widgetJSON, _ := json.Marshal(widget)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", widgetJSON)
}

// UpdateWidget updates widget
func (widgetController WidgetController) UpdateWidget(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Stub an widget to be populated from the body
	widget := models.Widget{}

	// Grab id
	vars := mux.Vars(r)
	id := vars["id"]

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Populate the widget data
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&widget)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Write the widget to mongo
	if err := widgetController.session.DB(common.AppConfig.Database).C("widgets").UpdateId(oid, &widget); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	widgetJSON, _ := json.Marshal(widget)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", widgetJSON)
}

// RemoveWidget removes an existing user resource
func (widgetController WidgetController) RemoveWidget(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

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

	// Remove widget
	if err := widgetController.session.DB(common.AppConfig.Database).C("widgets").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	// Write status
	w.WriteHeader(200)
}
