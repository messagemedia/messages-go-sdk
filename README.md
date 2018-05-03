# MessageMedia Messages Go SDK
[![Travis Build Status](https://api.travis-ci.org/messagemedia/messages-go-sdk.svg?branch=master)](https://travis-ci.org/messagemedia/messages-go-sdk)

The MessageMedia Messages API provides a number of endpoints for building powerful two-way messaging applications.

## ⭐️ Installing via Go Get
Install via Go Get:
run `go get github.com/messagemedia/messages-go-sdk`

## 🎬 Get Started
It's easy to get started. Simply enter the API Key and secret you obtained from the [MessageMedia Developers Portal](https://developers.messagemedia.com) into the code snippet below and a mobile number you wish to send to.

### 🚀 Send an SMS
* Destination numbers (`destination_number`) should be in the [E.164](http://en.wikipedia.org/wiki/E.164) format. For example, `+61491570156`.

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/messagemedia/messages-go-sdk"
    "github.com/messagemedia/messages-go-sdk/messages_pkg"
    "github.com/messagemedia/messages-go-sdk/models_pkg"
)

func main() {
    messagemedia_messages_sdk.Config.BasicAuthUserName = "YOUR_API_KEY"
    messagemedia_messages_sdk.Config.BasicAuthPassword = "YOUR_SECRET_KEY"

    messages := messages_pkg.NewMESSAGES()
    bodyValue := []byte(`{
        "messages": [{
            "content": "My first message",
            "destination_number": "YOUR_MOBILE_NUMBER",
            "format": "SMS"
        }]
    }`)

    var body *models_pkg.SendMessagesRequest
    json.Unmarshal(bodyValue, &body)

    result, err := messages.CreateSendMessages(body)
    fmt.Println(result, err)
}
```

### 🖼 Send an MMS
* Destination numbers (`destination_number`) should be in the [E.164](http://en.wikipedia.org/wiki/E.164) format. For example, `+61491570156`.

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/MessageMedia/messages-go-sdk"
    "github.com/MessageMedia/messages-go-sdk/messages_pkg"
    "github.com/MessageMedia/messages-go-sdk/models_pkg"
)

func main() {
    messagemedia_messages_sdk.Config.BasicAuthUserName = "YOUR_API_KEY"
    messagemedia_messages_sdk.Config.BasicAuthPassword = "YOUR_SECRET_KEY"

    messages := messages_pkg.NewMESSAGES()
    bodyValue := []byte(`{
       "messages":[
         {
            "content":"Test",
            "destination_number":"YOUR_MOBILE_NUMBER",
            "format": "MMS",
            "media":["https://upload.wikimedia.org/wikipedia/commons/6/6a/L80385-flash-superhero-logo-1544.png"]
         }
       ]
    }`)

    var body *models_pkg.SendMessagesRequest
    json.Unmarshal(bodyValue, &body)

    result, err := messages.CreateSendMessages(body)
    fmt.Println(result, err)
}
```

### 🕓 Get Status of a Message
You can get a messsage ID from a sent message by looking at the `message_id` from the response of the above example.
```go
package main

import (
    "fmt"
    "github.com/messagemedia/messages-go-sdk"
    "github.com/messagemedia/messages-go-sdk/messages_pkg"
)

func main() {
    messagemedia_messages_sdk.Config.BasicAuthUserName = "YOUR_API_KEY"
    messagemedia_messages_sdk.Config.BasicAuthPassword = "YOUR_SECRET_KEY"

    messages := messages_pkg.NewMESSAGES()
    messageId := "YOUR_MESSAGE_ID"

    var result interface{}
    result,_ = messages.GetMessageStatus(messageId)

    fmt.Println(result)
}

```

### 💬 Get replies to a message
You can check for replies that are sent to your messages
```go
package main

import (
    "fmt"
    "github.com/messagemedia/messages-go-sdk"
    "github.com/messagemedia/messages-go-sdk/replies_pkg"
    "github.com/messagemedia/messages-go-sdk/models_pkg"
)

func main() {
    messagemedia_messages_sdk.Config.BasicAuthUserName = "YOUR_API_KEY"
    messagemedia_messages_sdk.Config.BasicAuthPassword = "YOUR_API_SECRET"

    replies := replies_pkg.NewREPLIES()

    var result *models_pkg.CheckRepliesResponse
    result,_ = replies.GetCheckReplies()

    fmt.Println(result)
}

```

### ✅ Check Delivery Reports
This endpoint allows you to check for delivery reports to inbound and outbound messages.
```go
package main

import (
    "fmt"
    "github.com/messagemedia/messages-go-sdk"
    "github.com/messagemedia/messages-go-sdk/deliveryreports_pkg"
    "github.com/messagemedia/messages-go-sdk/models_pkg"
)

func main() {
    messagemedia_messages_sdk.Config.BasicAuthUserName = "YOUR_API_KEY"
    messagemedia_messages_sdk.Config.BasicAuthPassword = "YOUR_API_SECRET"

    deliveryReports := deliveryreports_pkg.NewDELIVERYREPORTS()

    var result *models_pkg.CheckDeliveryReportsResponse
    result,_ = deliveryReports.GetCheckDeliveryReports()

    fmt.Println(result)
}
```

## 📕 Documentation
Check out the [full API documentation](DOCUMENTATION.md) for more detailed information.

## 😕 Need help?
Please contact developer support at developers@messagemedia.com or check out the developer portal at [developers.messagemedia.com](https://developers.messagemedia.com/)

## 📃 License
Apache License. See the [LICENSE](LICENSE) file.
