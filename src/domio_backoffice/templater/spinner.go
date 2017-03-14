package templater

type TopBarData struct {
    LeftColumnLinks  []Link
    RightColumnLinks []Link
}

func GetSpinnerTemplate() string {
    return `{{ define "spinner_template"}}

                <div class="spinner-container">

                    <style>
                        .spinner {
                          position: absolute;
                          left: 50%;
                          margin-left: -25px;
                          top: 50%;
                          width: 50px;
                          text-align: center;
                          font-size: 10px;
                          height: 50px;
                          margin-top: -25px;
                          transition: opacity 0.3s;
                          opacity: 0;
                        }

                        .spinner.loading {
                            opacity: 1;
                        }

                        .spinner > div {
                            background-color: #158;
                            opacity: 0.8;
                            height: 100%;
                            width: 6px;
                            display: inline-block;
                        }

                        .spinner > div {
                           animation: sk-stretchdelay 1.2s infinite ease-in-out;
                        }

                        .spinner .rect2 {
                          animation-delay: -1.1s;
                        }

                        .spinner .rect3 {
                          animation-delay: -1.0s;
                        }

                        .spinner .rect4 {
                          animation-delay: -0.9s;
                        }

                        .spinner .rect5 {
                          animation-delay: -0.8s;
                        }

                        @-webkit-keyframes sk-stretchdelay {
                          0%, 40%, 100% { -webkit-transform: scaleY(0.4) }
                          20% { -webkit-transform: scaleY(1.0) }
                        }

                        @keyframes sk-stretchdelay {
                          0%, 40%, 100% {
                            transform: scaleY(0.4);
                            -webkit-transform: scaleY(0.4);
                          }  20% {
                            transform: scaleY(1.0);
                            -webkit-transform: scaleY(1.0);
                          }
                        }
                    </style>

                    <div class="spinner loading">
                      <div class="rect1"></div>
                      <div class="rect2"></div>
                      <div class="rect3"></div>
                      <div class="rect4"></div>
                      <div class="rect5"></div>
                    </div>

                </div>

            {{end}}
        `
}

func GetTopBarData(pageName string, userName string, token string) TopBarData {
    var customLinks []Link

    if (userName != "") {

        customLinks = []Link{
            {Url:"/profile", Label:userName},
            {Url:"/logout", Label:"Logout"},
        }
    } else {
        customLinks = []Link{
            {Url:"/login", Label:"Login"},
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