package server

import (
    "domio_backoffice/router"
    "log"
    "domio_backoffice/components/config"
    "net/http"
    "fmt"
    "os"
)

func StartRouter() {
    log.Print("Starting router...")
    domiorouter := router.NewRouter()
    conf := config.Config

    if (conf.Env == "development") {
        log.Print("Development environment, handling static files by Go...")
        domiorouter.Path("/style.css").Handler(http.FileServer(http.Dir("/usr/local/domio_backoffice_client")))
        domiorouter.Path("/app/app.js").Handler(http.FileServer(http.Dir("/usr/local/domio_backoffice_client")))

        windows_dart_path := "/Users/sbasharov/WebstormProjects/DomioClient/src"
        mac_dart_path := "/Users/sergeibasharov/WebstormProjects/DomioClient/src"

        current_dart_path := ""

        _, err := os.Stat(windows_dart_path)

        if err == nil {
            current_dart_path = windows_dart_path
        } else {
            current_dart_path = mac_dart_path
        }
        domiorouter.PathPrefix("/app").Handler(http.FileServer(http.Dir(current_dart_path)))
    }

    log.Printf("Web server is running on http://localhost:%d", config.Config.Port)
    err := http.ListenAndServe(fmt.Sprintf(":%v", config.Config.Port), domiorouter)
    if (err != nil) {
        msg := fmt.Sprintf("Failed to start web server on port %d", config.Config.Port)
        log.Fatal(msg)
    }

}
func Start() {
    fmt.Print("Starting app...\n")
    StartRouter()
}
