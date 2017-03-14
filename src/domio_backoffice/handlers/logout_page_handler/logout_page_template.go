package logout_page_handler

import (
    "html/template"
    "log"
)

func GetLogoutPageTemplate(parsedTemplate *template.Template) {
    _, err := parsedTemplate.New("main_template").Parse(`{{define "main_template"}}
                                                            <p>You have been logged out.</p>
                                                            <p><a href='/'>Home</a></p>
                                                         {{end}}`)
    if (err != nil) {
        log.Print(err)
    }

}
