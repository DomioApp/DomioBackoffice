package login_page_handler

import (
    "html/template"
    "net/http"
    "domio_backoffice/templater"
)

type FormRow struct {
    Label string
    Name  string
    Type  string
}

type SubmitButton struct {
    Label string
}

type LoginPageData struct {
    FormRows     []FormRow
    SubmitButton SubmitButton
}

type PageData struct {
    PageTitle     string
    LoginPageData LoginPageData
}

var loginPageTemplate *template.Template

func init() {
    loginPageTemplate = templater.BuildTemplate(GetLoginPageTemplate)
}

func LoginPageHandler(w http.ResponseWriter, req *http.Request) {

    templater.WriteTemplate(w, req, loginPageTemplate, GetPageName(), GetPageData())

}

func GetUrl() string {
    return "/login"
}

func GetPageName() string {
    return "LoginPage"
}

func GetPageData() PageData {

    pageData := PageData{
        PageTitle: "Domio - Marketplace for domains. Login to your account.",

        LoginPageData:LoginPageData{
            FormRows:[]FormRow{
                {Label:"Email", Name:"email", Type:"email"},
                {Label:"Password", Name:"password", Type:"password"},
            },
            SubmitButton: SubmitButton{Label:"Submit"},
        },
    }
    return pageData
}
