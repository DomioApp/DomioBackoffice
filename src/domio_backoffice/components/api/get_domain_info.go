package api

import (
    "domio_backoffice/errors"
    "fmt"
    "net/http"
)

func GetDomainInfo(domainName string, token string) (*DomainJson, *errors.DomioError) {

    var domain DomainJson

    req := Request{
        Url:fmt.Sprintf("/domain/%s", domainName),
        Method:http.MethodGet,
        Token:token,
    }

    //if (resp.StatusCode == http.StatusNotFound) {
    //    return nil, &errors.DomainNotFound
    //}

    req.Run(&domain)

    return &domain, nil
}
