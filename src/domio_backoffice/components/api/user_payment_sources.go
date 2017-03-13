package api

import (
    "domio_backoffice/components/config"
    "net/http"
    "time"
    "log"
    "fmt"
)

type PaymentSource struct {
    Id       string `json:"id"`
    ExpMonth uint64 `json:"exp_month"`
    ExpYear  uint64 `json:"exp_year"`
    LastFour string `json:"last4"`
    Brand    string `json:"brand"`
    Currency string `json:"currency"`
    Country  string `json:"country"`
    Name     string `json:"name"`
}

func GetUserPaymentSources(token string) []PaymentSource {

    var paymentSources []PaymentSource

    url := config.Config.ApiUrl

    req, requestErr := http.NewRequest("GET", url + "/cards", nil)

    req.Header.Set("Content-Type", "application/json")

    timeout := time.Duration(10 * time.Second)

    client := &http.Client{
        Timeout: timeout,
    }

    req.Header.Add("Authorization", "Bearer " + token)

    resp, requestErr := client.Do(req)

    if requestErr != nil {
        log.Print(requestErr)
        return paymentSources
    }
    defer resp.Body.Close()

    decodeError := DecodeJsonRequestBody(resp, &paymentSources)

    if (decodeError != nil) {
        log.Print(decodeError)
        return paymentSources
    }

    return paymentSources
}

func GetUserPaymentSource(token string, paymentSourceId string) PaymentSource {

    var paymentSource PaymentSource

    url := config.Config.ApiUrl

    req, requestErr := http.NewRequest("GET", fmt.Sprintf("%s/cards/%s", url, paymentSourceId), nil)

    req.Header.Set("Content-Type", "application/json")

    timeout := time.Duration(10 * time.Second)

    client := &http.Client{
        Timeout: timeout,
    }

    req.Header.Add("Authorization", "Bearer " + token)

    resp, requestErr := client.Do(req)

    if requestErr != nil {
        log.Print(requestErr)
        return paymentSource
    }
    defer resp.Body.Close()

    decodeError := DecodeJsonRequestBody(resp, &paymentSource)

    if (decodeError != nil) {
        log.Print(decodeError)
        return paymentSource
    }

    return paymentSource
}
