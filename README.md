# slacknotifications

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

# Objective

Slack allows you to create webhooks which can be used as an API endpoints and sent messages programmatically.
A many time, we do want our process to send us the slack notifications.

### Prerequisites
There are no dependencies on any of the GoLang packages.

### Imprting Package

```
import (
  "github.com/nitor-infotech-oss/go-slacknotification/slacknotifications"
)
```

Package can be imported by simply using the path of repo where it resides.

### Example Usage

```
package main
import (
  "fmt"
  "github.com/nitor-infotech-oss/go-slacknotification/slacknotifications"
)

func main() {
  webhookUrl := "https://hooks.slack.com/services/TKDC1KVV1/XXXXXXXXXXXXXXX/XXXXXXXXXXXXXXX"
  err = slacknotifications.SendSlackNotification(webhookUrl, summaryString)
  if err != nil {
    fmt.Println("Error while sending msg to slack")
  }
}
```

 ### Steps to generate webhook url from slack

- To get the correct webhook url on which we have post the messages, need to add configuration on the incoming WebHoks page.

- Go to https://my.slack.com/services/new/incoming-webhook/ link, it may ask you to log in
![alt text](https://i.ibb.co/g7m9vmY/webhook1.png)
- Choose the channel you want to post the messages and click on the Add Incoming Webhooks integration button.
- Once you click above mentioned button this will give you the webhook url (e.g. https://hooks.slack.com/services/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)
![alt text](https://i.ibb.co/M1TYCNW/webhook2.png)
