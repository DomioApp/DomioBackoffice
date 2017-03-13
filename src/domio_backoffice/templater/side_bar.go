package templater

type SideBarData struct {
    Title string
    Links []Link
}

func GetSideBarTemplate() string {

    sideBarTemplate := `
                         {{define "sidebar_template"}}
                             {{with .PageData.SideBarData}}

                                <div class="b-side-bar-container">
                                    <h4>{{.Title}}</h4>

                                    {{range .Links}}
                                         <a href="{{.Url}}">{{.Label}}</a>
                                    {{end}}
                                </div>

                             {{end}}
                         {{end}}
                        `
    return sideBarTemplate

}

func GetSideBarData() SideBarData {
    sideBarData := SideBarData{
        Title: "Categories",
        Links:[]Link{
            {Url:"/domains/business", Label:"Business"},
            {Url:"/domains/social", Label:"Social"},
            {Url:"/domains/cars", Label:"Cars"},
        },
    }

    return sideBarData
}