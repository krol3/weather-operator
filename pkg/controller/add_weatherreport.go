package controller

import (
	"github.com/krol3/weather-operator/pkg/controller/weatherreport"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, weatherreport.Add)
}
