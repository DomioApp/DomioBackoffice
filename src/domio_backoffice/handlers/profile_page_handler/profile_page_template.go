package profile_page_handler

import (
    "html/template"
    "log"
)

func GetProfilePageTemplate(parsedTemplate *template.Template) {
    _, err := parsedTemplate.New("main_template").Parse(`{{define "main_template"}}
                                                            {{template "profile_top_bar_template" .}}
                                                            {{template "profile_main_area_template" .}}
                                                         {{end}}`)
    if (err != nil) {
        log.Print(err)
    }

    _, err2 := parsedTemplate.New("profile_top_bar_template").Parse(GetProfileTopBarTemplate())

    if (err2 != nil) {
        log.Print(err2)
    }

    _, err3 := parsedTemplate.New("profile_main_area_template").Parse(GetProfileMainAreaTemplate())

    if (err3 != nil) {
        log.Print(err3)
    }
}
