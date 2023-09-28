package controllers

import (
	"net/http"

	"github.com/vrtttx/gocoffee/helpers"
	"github.com/vrtttx/gocoffee/services"
)

func GetAllCoffees(w http.ResponseWriter, r *http.Request) {
	var coffees services.Coffee

	all, err := coffees.GetAllCoffees()

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)

		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"coffees": all})
}