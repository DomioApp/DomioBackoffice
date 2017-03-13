package templater

func GetAppStatusInfoBarTemplate() string {
    appStatusInfoBarTemplate := `{{ define "app_status_infobar"}}
                                    {{with .AppStatusInfoBarData}}
                                        <div class="b-app-status-info">
                                            <!--<div><span>API Info</span></div>-->
                                            <div>API Version: <span>{{.Version}}</span></div>
                                            <div>Built: <span>{{.BuildAgo}} ago</span></div>
                                            <!--<div>Buildstamp: <span>{{.Buildstamp}}</span></div>-->
                                            <!--<div>BuildDateTime: <span>{{.BuildDateTime}}</span></div>-->
                                            <!--<div>Hash: <span>{{.Hash}}</span></div>-->
                                        </div>
                                    {{end}}
                                 {{end}}
                                `
    return appStatusInfoBarTemplate
}
