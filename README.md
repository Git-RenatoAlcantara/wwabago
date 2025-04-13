# Go WhatsApp API Library

A lightweight and easy-to-use Go library for sending and receiving messages using the **official WhatsApp Business Cloud API**.

## ✨ Features

- Send text and media messages via WhatsApp
- Receive and handle incoming messages via webhooks
- Fully compatible with the official WhatsApp Business API
- Built for simplicity and performance in Go applications

## 📦 Installation

```bash
go get github.com/your-username/your-whatsapp-lib
```

```go
package main

import (
    "log"
    "your-whatsapp-lib"
)

func main() {
    client := whatsapp.NewClient("your_access_token", "your_phone_number_id")
    
    err := client.SendTextMessage("recipient_number", "Hello from Go!")
    if err != nil {
        log.Fatal(err)
    }
}
```



