package templater

func getBaseTemplateContent() string {
    baseTemplateContent := `
                        {{define "base_template"}}
                            <!DOCTYPE html>

                            <html lang="en">

                            <head>
                                <meta charset="UTF-8">
                                <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0"/>
                                <meta name="page" content="{{.BaseTemplateData.PageName}}">

                                <title>{{.PageData.PageTitle}}</title>
                                <link rel="stylesheet" href="/style.css" />
                            </head>

                            <body>

                                {{if eq .BaseTemplateData.PageName "LoginPage"}}

                                    {{template "main_template" .}}

                                {{else}}

                                    {{template "spinner_template" .}}

                                {{end}}


                                {{if eq .UseDart true}}

                                    {{if eq .BaseTemplateData.PageName "LoginPage"}}
                                        <script type="application/dart" src="/app/login_page.dart"></script>
                                    {{else}}
                                        <script type="application/dart" src="/app/app.dart"></script>
                                    {{end}}

                                {{else}}

                                    {{if eq .BaseTemplateData.PageName "LoginPage"}}
                                        <script type="application/javascript" src="/app/login_page.js"></script>
                                    {{else}}
                                        <script type="application/javascript" src="/app/app.js"></script>
                                    {{end}}

                                {{end}}

                            </body>

                            </html>
                        {{end}}`

    return baseTemplateContent
}