package controllers

import (
	"aahframework.org/aah.v0"
)

// AppController struct application controller
type AppController struct {
	*aah.Context
}

// Before method is called for all the application requests
// before the controllers action gets called.
func (a *AppController) Before() {
	releases, _ := aah.AppConfig().StringList("docs.releases")
	a.AddViewArg("Releases", releases)
}
