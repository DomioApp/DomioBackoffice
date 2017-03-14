package dashboard_page_handler

import (
    "html/template"
    "net/http"
    "domio_backoffice/templater"
    "domio_backoffice/components/api"
    "log"
)

type PageData struct {
    PageTitle           string
    ProfileTopBarData   ProfileTopBarData
    DashboardMainAreaData DashboardMainAreaData
}

var profilePageTemplate *template.Template
var tokenCookie *http.Cookie

func init() {
    profilePageTemplate = templater.BuildTemplate(GetProfilePageTemplate)
}

func ProfilePageHandler(w http.ResponseWriter, req *http.Request) {
    var err error
    tokenCookie, err = req.Cookie("token")

    if (err != nil) {
        http.Redirect(w, req, "/login", http.StatusTemporaryRedirect)

    } else {
        templater.WriteTemplate(w, req, profilePageTemplate, GetPageName(), GetPageData())
    }

}

func GetUrl() string {
    return "/dashboard"
}

func GetPageName() string {
    return "DashboardPage"
}

func GetPageData() PageData {
    return PageData{
        PageTitle: "Domio Backoffice Dashboard",
        ProfileTopBarData:GetProfileTopBarData(),
        DashboardMainAreaData:DashboardMainAreaData{Title:"Dashboard"},
    }

}

func GetProfileTopBarData() ProfileTopBarData {

    countResp := api.GetUserDomainsCount(tokenCookie.Value)

    log.Print(countResp)

    return ProfileTopBarData{
        Links:[]templater.Link{
            {Url:"/subscriptions", Label:"Subscriptions"},
            {Url:"/domains", Label:"All Domains"},
            {Url:"/domains/pending", Label:"Pending Domains"},
        },
    }
}