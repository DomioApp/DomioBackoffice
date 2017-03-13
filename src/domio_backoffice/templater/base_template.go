package templater

func getBaseTemplateContent() string {
    baseTemplateContent := `
                        {{define "base_template"}}
                            <!DOCTYPE html>

                            <html lang="en">

                            <head>
                                <meta charset="UTF-8">
                                <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0"/>
                                <meta name="description" content="Domio is a marketplace for domains.">
                                <meta name="page" content="{{.BaseTemplateData.PageName}}">

                                <title>{{.PageData.PageTitle}}</title>
                                <link rel="stylesheet" href="/style.css" />
                            </head>

                            <body>
                                {{template "top_bar_template" .}}
                                {{template "main_template" .}}
                                {{template "app_status_infobar" .}}

                                {{if eq .UseDart true}}
                                    <script type="application/dart" src="/app/app.dart"></script>
                                {{else}}
                                    <script src="/app/app.js"></script>
                                {{end}}

                            </body>

                            </html>
                        {{end}}`

    return baseTemplateContent
}