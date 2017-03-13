package login_page_handler

import (
    "html/template"
    "log"
    "domio_backoffice/templater"
)

func GetLoginPageTemplate(parsedTemplate *template.Template) {
    _, err := parsedTemplate.New("main_template").Parse(`{{define "main_template"}}
                                                            {{template "login_form_template" .}}
                                                         {{end}}`)
    if (err != nil) {
        log.Print(err)
    }

    _, err2 := parsedTemplate.New("sidebar_template").Parse(templater.GetSideBarTemplate())

    if (err2 != nil) {
        log.Print(err2)
    }

    _, err3 := parsedTemplate.New("login_form_template").Parse(GetLoginFormTemplate())

    if (err3 != nil) {
        log.Print(err3)
    }
}
