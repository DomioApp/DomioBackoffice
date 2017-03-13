package api

import (
    "net/http"
)

type DomainsCountResponse struct {
    Count int `json:"count"`
}

func GetUserDomainsCount(token string) DomainsCountResponse {

    var userDomainsCount DomainsCountResponse

    req := Request{
        Url:"/user/domains/count",
        Method:http.MethodGet,
        Token:token,
    }

    req.Run(&userDomainsCount)

    return userDomainsCount

}
