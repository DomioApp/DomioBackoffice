package api

import (
    "net/http"
)

func GetDomains() []DomainJson {

    var domains []DomainJson

    req := Request{
        Url:"/domains/available",
        Method:http.MethodGet,
    }

    req.Run(&domains)

    return domains
}
