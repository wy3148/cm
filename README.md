# Simple github api front end
    The project is implemented in Go programming language, we use the opensource
    "github.com/spf13/cobra" as the cli/console application framework 

# How to build
    install Go, https://golang.org/doc/install
    make sure $GOPATH has been configured

```
    cd $GOPATH
    go get github.com/wy3148/cm
    cd $GOPATH/src/github.com/wy3148/cm
    go build
```

# How to use
    ./cm help
    ./cm client --api_key your_api_key --client_id your_client_id
    ./cm list create --api_key your_api_key --client_id your_client_id --json_file example.json

    example.json is a json-formatted file
    {
    "Title": "Website Subscribersi Testing",
    "UnsubscribePage": "http://www.example.com/unsubscribed.html",
    "UnsubscribeSetting": "AllClientLists",
    "ConfirmedOptIn": false,
    "ConfirmationSuccessPage": "http://www.example.com/joined.html"
    }

# API url
    the default API url: https://api.createsend.com/api/v3.1
    run 'export cm_url' on linux/mac or 'set cm_url' on windows to configure a different API url

# Testing result
./cm client --api_key your_api_key --client_id your_client_id
using default url: https://api.createsend.com/api/v3.1
2018/02/17 22:23:41 response code: 200
2018/02/17 22:23:41 body response:
```
 {
 "ApiKey": "28160ca6ab63eda45a74d3317869b500b364b85fc31ead15",
 "BasicDetails": {
  "ClientID": "e7a8cbe95cd62de2afda7189612a5895",
  "CompanyName": "NETFON",
  "ContactName": "",
  "EmailAddress": "",
  "Country": "Australia",
  "TimeZone": "(GMT+10:00) Canberra, Melbourne, Sydney"
 },
 "BillingDetails": {
  "CanPurchaseCredits": false,
  "Credits": 0,
  "BaseDeliveryRate": 7.0,
  "BaseRatePerRecipient": 1.4000,
  "BaseDesignSpamTestRate": 7.0,
  "MarkupOnDelivery": 0.0,
  "MarkupPerRecipient": 0.0,
  "MarkupOnDesignSpamTest": 0.0,
  "Currency": "AUD",
  "ClientPays": false
 }
}
```



./cm list create --api_key your_api_key --client_id your_client_id --json_file example.json
using default url: https://api.createsend.com/api/v3.1

```
2018/02/17 22:21:50 will create a new list: {
    "Title": "Website Subscribersi Testing",
    "UnsubscribePage": "http://www.example.com/unsubscribed.html",
    "UnsubscribeSetting": "AllClientLists",
    "ConfirmedOptIn": false,
    "ConfirmationSuccessPage": "http://www.example.com/joined.html"
}

2018/02/17 22:21:51 response code: 201
2018/02/17 22:21:51 body response:
 "aaf13cd89855e2c164350cbe4cd4dba7"

```

