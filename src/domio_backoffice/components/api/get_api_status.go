package api

import (
    "net/http"
)

func GetAPIStatus() *AppStatusInfo {

    var appStatusInfo AppStatusInfo

    req := Request{
        Url:"",
        Method:http.MethodGet,

    }

    req.Run(&appStatusInfo)

    return &appStatusInfo
}
