package config

import (
    "encoding/json"
    "os"
    "log"
    "strconv"
    "time"
    "path"
)

type Configuration struct {
    UseDart bool `json:"use_dart"`
    Port    uint    `json:"port"`
    Env     string    `json:"env"`
    ApiUrl  string    `json:"api_url"`
}

type AppStatus struct {
    Buildstamp string `json:"app_buildstamp"`
    Hash       string `json:"app_hash"`
    Version    string `json:"app_version"`
}

func (*AppStatus) GetBuildAgoValue() string {

    i, err := strconv.ParseInt(AppStatusInfo.Buildstamp, 10, 64)
    if err != nil {
        panic(err)
    }
    tm := time.Unix(i, 0)

    return time.Since(tm).String()
}

func (*AppStatus) GetBuildDateTime() string {

    i, err := strconv.ParseInt(AppStatusInfo.Buildstamp, 10, 64)
    if err != nil {
        panic(err)
    }
    tm := time.Unix(i, 0)

    return tm.String()
}

var AppStatusInfo AppStatus
var Config Configuration
var ConfigPath = "/usr/local/domio_public"

func LoadConfig() error {
    configFile := path.Join(ConfigPath, "config.json")
    if _, err := os.Stat(configFile); os.IsNotExist(err) {
        log.Fatalln("error:", err)
    }

    file, _ := os.Open(configFile)

    decoder := json.NewDecoder(file)
    err := decoder.Decode(&Config)

    if err != nil {
        log.Fatalln("Config file load error:", err)
    }

    return nil
}