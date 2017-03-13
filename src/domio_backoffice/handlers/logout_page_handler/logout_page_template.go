package logout_page_handler

import (
    "html/template"
    "log"
    "domio_backoffice/templater"
)

func GetLogoutPageTemplate(parsedTemplate *template.Template) {
    _, err := parsedTemplate.New("main_template").Parse(`{{define "main_template"}}
                                                            <p>You have been logged out.</p>
                                                            <p><a href='/'>Home</a></p>
                                                         {{end}}`)
    if (err != nil) {
        log.Print(err)
    }

    _, err2 := parsedTemplate.New("sidebar_template").Parse(templater.GetSideBarTemplate())

    if (err2 != nil) {
        log.Print(err2)
    }
}
