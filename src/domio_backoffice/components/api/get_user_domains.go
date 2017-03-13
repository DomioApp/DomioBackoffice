package api

import (
    "net/http"
)

func GetUserDomains(token string) []DomainJson {

    var userDomains []DomainJson

    req := Request{
        Url:"/user/domains",
        Method:http.MethodGet,
        Token:token,
    }

    req.Run(&userDomains)

    return userDomains

}
