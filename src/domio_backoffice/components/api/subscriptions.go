package api

import (
    "domio_backoffice/components/config"
    "net/http"
    "time"
    "log"
)

func GetSubscription(subscriptionId string, token string) Subscription {

    var subscription Subscription

    url := config.Config.ApiUrl

    req, requestErr := http.NewRequest("GET", url + "/subscriptions/" + subscriptionId, nil)

    req.Header.Set("Content-Type", "application/json")

    timeout := time.Duration(5 * time.Second)

    client := &http.Client{
        Timeout: timeout,
    }

    req.Header.Add("Authorization", "Bearer " + token)

    resp, requestErr := client.Do(req)

    if requestErr != nil {
        log.Print(requestErr)
        return subscription
    }
    defer resp.Body.Close()

    decodeError := DecodeJsonRequestBody(resp, &subscription)

    if (decodeError != nil) {
        log.Print(decodeError)
        return subscription
    }

    return subscription
}
