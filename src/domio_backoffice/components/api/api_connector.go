package api

type DomainJson struct {
    Name          string `json:"name"`
    Owner         string `json:"owner"`
    PricePerMonth uint64 `json:"price_per_month"`
    IsVisible     bool `json:"is_visible"`
    IsRented      bool `json:"is_rented"`
    ZoneId        string `json:"zone_id"`
    NS1           string `json:"ns1"`
    NS2           string `json:"ns2"`
    NS3           string `json:"ns3"`
    NS4           string `json:"ns4"`
}

type Plan struct {
    Id       string `json:"id"`
    Amount   uint64 `json:"amount"`
    Created  uint64 `json:"created"`
    Currency string `json:"currency"`
    Interval string `json:"interval"`
}

type SubMetadata struct {
    Domain string `json:"domain"`
}

type Subscription struct {
    Id       string `json:"id"`
    Quantity uint64 `json:"quantity"`
    Status   string `json:"status"`
    Created  uint64 `json:"created"`
    Start    uint64 `json:"start"`

    Plan     Plan
    Metadata SubMetadata
}

type AppStatusInfo struct {
    Version       string `json:"app_version"`
    BuildAgo      string `json:"app_buildago"`
    Buildstamp    string `json:"app_buildstamp"`
    BuildDateTime string `json:"app_builddatetime"`
    Hash          string `json:"app_hash"`
}
