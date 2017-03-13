package api

import (
    "domio_backoffice/components/config"
    "net/http"
    "time"
    "log"
    "fmt"
)

type RecordValue struct {
    Value string
}

type GeoLocation struct {
    ContinentCode   string
    CountryCode     string
    SubdivisionCode string
}

type Record struct {
    Name                    string
    AliasTarget             string
    Failover                string
    GeoLocation             GeoLocation
    HealthCheckId           string
    Region                  string
    TrafficPolicyInstanceId string
    Weight                  uint64
    Type                    string
    SetIdentifier           string
    ResourceRecords         []RecordValue
    TTL                     int64
}

func GetSubscriptionRecords(subscriptionId string, token string) []Record {

    var subRecords []Record

    url := config.Config.ApiUrl

    req, requestErr := http.NewRequest("GET", fmt.Sprintf("%s/subscriptions/%s/records", url, subscriptionId), nil)

    req.Header.Set("Content-Type", "application/json")

    timeout := time.Duration(5 * time.Second)

    client := &http.Client{
        Timeout: timeout,
    }

    req.Header.Add("Authorization", "Bearer " + token)

    resp, requestErr := client.Do(req)

    if requestErr != nil {
        log.Print(requestErr)
        return subRecords
    }
    defer resp.Body.Close()

    decodeError := DecodeJsonRequestBody(resp, &subRecords)

    if (decodeError != nil) {
        log.Print(decodeError)
        return subRecords
    }

    return subRecords
}
