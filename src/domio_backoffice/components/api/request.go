package api

import (
    "net/http"
    "time"
    "log"
    "domio_backoffice/errors"
    "fmt"
    "domio_backoffice/components/config"
)

type Request struct {
    Url    string
    Method string
    Token  string
}

func (request *Request) Run(responseObject interface{}) *errors.DomioError {

    req, requestErr := http.NewRequest(request.Method, fmt.Sprintf("%s%s", config.Config.ApiUrl, request.Url), nil)

    req.Header.Set("Content-Type", "application/json")

    if (request.Token != "") {
        req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", request.Token))
    }

    timeout := time.Duration(5 * time.Second)

    client := &http.Client{
        Timeout: timeout,
    }

    resp, requestErr := client.Do(req)

    if requestErr != nil {
        log.Print(requestErr)
        return &errors.UnknownError
    }

    decodeError := DecodeJsonRequestBody(resp, &responseObject)

    if (decodeError != nil) {
        log.Print(decodeError)
        return &errors.UnknownError
    }

    defer resp.Body.Close()

    return nil
}