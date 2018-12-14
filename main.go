package main

import "os"
import "fmt"
import "encoding/json"

type Config struct {
    Database struct {
        Host     string
    }
    Port int
}

func LoadConfiguration(file string) Config {
    var config Config
    configFile, err := os.Open(file)
    defer configFile.Close()
    if err != nil {
        fmt.Println(err.Error())
    }
    jsonParser := json.NewDecoder(configFile)
    jsonParser.Decode(&config)
    return config
}

func main() {
    config := LoadConfiguration("config.json")
    fmt.Println(config)
    fmt.Println(config.Database.Host)
    fmt.Println(config.Port)
}