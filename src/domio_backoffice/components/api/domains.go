package api

import (
    "domio_backoffice/components/config"
    "net/http"
    "time"
    "log"
)

func GetSubscriptions(token string) []Subscription {

    var subscriptions []Subscription

    url := config.Config.ApiUrl

    req, requestErr := http.NewRequest("GET", url + "/subscriptions", nil)

    req.Header.Set("Content-Type", "application/json")

    timeout := time.Duration(5 * time.Second)

    client := &http.Client{
        Timeout: timeout,
    }

    req.Header.Add("Authorization", "Bearer " + token)

    resp, requestErr := client.Do(req)

    if requestErr != nil {
        log.Print(requestErr)
        return subscriptions
    }
    defer resp.Body.Close()

    decodeError := DecodeJsonRequestBody(resp, &subscriptions)

    if (decodeError != nil) {
        log.Print(decodeError)
        return subscriptions
    }

    return subscriptions
}
