package profile_page_handler

import "domio_backoffice/templater"

type ProfileTopBarData struct {
    Links []templater.Link
}

func GetProfileTopBarTemplate() string {
    return `{{ define "profile_top_bar_template"}}

                {{with .PageData.ProfileTopBarData}}

                    <div class="b-profile-top-bar-container">
                        {{range .Links}}
                             <a href="{{.Url}}">{{.Label}}</a>
                        {{end}}
                    </div>

                {{end}}

            {{end}}
        `
}