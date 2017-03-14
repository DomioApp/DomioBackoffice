package login_page_handler

import (
    "html/template"
    "log"
)

func GetLoginPageTemplate(parsedTemplate *template.Template) {
    _, err := parsedTemplate.New("main_template").Parse(`{{define "main_template"}}
                                                            {{template "login_form_template" .}}
                                                         {{end}}`)
    if (err != nil) {
        log.Print(err)
    }

    _, err3 := parsedTemplate.New("login_form_template").Parse(GetLoginFormTemplate())

    if (err3 != nil) {
        log.Print(err3)
    }
}
