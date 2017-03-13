package dashboard_page_handler

type DashboardMainAreaData struct {
    Title string
}

func GetProfileMainAreaTemplate() string {
    return `{{ define "profile_main_area_template"}}

                {{with .PageData.DashboardMainAreaData}}

                    <div class="b-profile-main-area-container">

                        <input type="checkbox"/> Some option<br>
                        <input type="checkbox"/> Some option 1<br>
                        <hr>

                        <button class="b-delete-account-button">Delete my account</button>
                    </div>

                {{end}}

            {{end}}
        `
}