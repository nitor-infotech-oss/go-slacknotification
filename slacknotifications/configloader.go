package slacknotifications

import (
  "encoding/json"
  "fmt"
  "os"
)

// struct object to load slack configurations
type SlackConfig struct {
  WebookURL     string `json:"webhookurl"`
}

// method that reads config file and converts it to Config strut
func LoadConfig(file string) SlackConfig {
  var config SlackConfig
  configFile, err := os.Open(file)
  defer configFile.Close()
  if err != nil {
      fmt.Println(err.Error())
  }
  jsonParser := json.NewDecoder(configFile)
  jsonParser.Decode(&config)
  return config
}
