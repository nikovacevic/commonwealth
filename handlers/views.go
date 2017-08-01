package handlers

import "github.com/nikovacevic/commonwealth/views"

var adminView = views.NewView("default", "views/admin/index.gohtml")
var error404View = views.NewView("default", "views/404.gohtml")
var indexView = views.NewView("default", "views/index.gohtml")
var loginView = views.NewView("default", "views/auth/login.gohtml")
var registerView = views.NewView("default", "views/auth/register.gohtml")
var usersView = views.NewView("default", "views/users/index.gohtml")
