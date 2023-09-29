package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vrtttx/gocoffee/helpers"
	"github.com/vrtttx/gocoffee/services"
)

var models services.Models
var coffee = models.Coffee

func GetAllCoffees(w http.ResponseWriter, r *http.Request) {
	var coffees services.Coffee

	all, err := coffees.GetAllCoffees()

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)

		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"coffees": all})
}

func GetCoffeeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	coffee, err := coffee.GetCoffeeById(id)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)

		return
	}

	helpers.WriteJSON(w, http.StatusOK, coffee)
}