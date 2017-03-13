package api

import (
    "net/http"
    "encoding/json"
    "domio_backoffice/errors"
    "log"
)

func DecodeJsonRequestBody(resp *http.Response, obj interface{}) *errors.DomioError {

    decoder := json.NewDecoder(resp.Body)

    decodeError := decoder.Decode(&obj)
    if (decodeError != nil) {
        log.Print("#######################################")
        log.Print(decodeError)
        log.Print("#######################################")
        return &errors.JsonDecodeError
    }

    return nil
}