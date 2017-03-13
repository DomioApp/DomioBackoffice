package templater

type TopBarData struct {
    LeftColumnLinks  []Link
    RightColumnLinks []Link
}

func GetTopBarTemplate() string {
    return `{{ define "top_bar_template"}}

                {{with .TopBarData}}

                    <div class="b-top-bar-container">
                        <div class="toggle-menu-button"></div>
                        <div class="left-area">
                            {{range .LeftColumnLinks}}
                                 <a {{if .ClassName}}class="{{.ClassName}}"{{end}} href="{{.Url}}">{{.Label}}</a>
                            {{end}}
                        </div>

                        <div class="right-area">
                            {{range .RightColumnLinks}}
                                 <a {{if .ClassName}}class="{{.ClassName}}"{{end}} href="{{.Url}}">{{.Label}}</a>
                            {{end}}
                        </div>

                    </div>

                {{end}}

            {{end}}
        `
}

func GetTopBarData(pageName string, userName string, token string) TopBarData {
    var customLinks []Link

    if (userName != "") {

        //countResp := api.GetUserDomainsCount(token)

        customLinks = []Link{
            //{Url:"/profile/domains", Label:fmt.Sprintf("My Domains (%d)", countResp.Count)},
            {Url:"/profile", Label:userName},
            {Url:"/logout", Label:"Logout"},
        }
    } else {
        customLinks = []Link{
            {Url:"/profile/domains/add", Label:"Add Domain", ClassName:"b-top-bar-container__domain-add-link"},
            {Url:"/login", Label:"Login"},
            {Url:"/signup", Label:"Signup"},
        }
    }

    dataset := TopBarData{
        LeftColumnLinks:[]Link{
            {Url:"/dashboard", Label:"Dashboard"},
            {Url:"/users", Label:"Users"},
        },
        RightColumnLinks:customLinks,
    }

    return dataset
}