package routes

import (
    "net/http"
    "domio_backoffice/handlers/login_page_handler"
    "domio_backoffice/handlers/profile_page_handler"
    "domio_backoffice/handlers/logout_page_handler"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var RoutesList = Routes{
    Route{
        "LoginPage",
        "GET",
        login_page_handler.GetUrl(),
        login_page_handler.LoginPageHandler,
    },
    Route{
        "LogoutPage",
        "GET",
        logout_page_handler.GetUrl(),
        logout_page_handler.LogoutPageHandler,
    },
    Route{
        "ProfilePage",
        "GET",
        profile_page_handler.GetUrl(),
        profile_page_handler.ProfilePageHandler,
    },
}

