
package main

import (
  "fmt"
  "github.com/nitor-infotech-oss/go-slacknotification/slacknotifications"
)

func main() {
  configFilename = "config.json"
  msg := "This is sample slack notification." 
  err = slacknotifications.SendSlackNotification(configFilename, msg)
  if err != nil {
    fmt.Println("Error while sending msg to slack")
  }
}
