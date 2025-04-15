# Go WhatsApp API Library

A lightweight and easy-to-use Go library for sending and receiving messages using the **official WhatsApp Business Cloud API**.

## ✨ Features

- Send text and media messages via WhatsApp
- Receive and handle incoming messages via webhooks
- Fully compatible with the official WhatsApp Business API
- Built for simplicity and performance in Go applications

## 📦 Installation

```bash
	go get github.com/Git-RenatoAlcantara/wwabago
```

## Send text message
```go
package main

import (
	"fmt"
	"log"

	wwabago "github.com/Git-RenatoAlcantara/wwabago"
)


func main(){
	wwaba , err := wwabago.CreateWwaba("access_token", "identification_number_id")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	msg := wwabago.NewMessage("+00(00)00000-0000", "👋 Hello! I'm using the WWabago library.")
	_, err = wwaba.Send(ctx, msg)
	fmt.Println(err) // ou log.Println(err)

}
```
## Send image message

```go
package main

import (
	"fmt"
	"log"

	wwabago "github.com/Git-RenatoAlcantara/wwabago"
)


func main(){

	wwaba , err := wwabago.CreateWwaba("access_token", "identification_number_id")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	msg := wwabago.NewImageMessage(
		"+00(00)00000-0000",
		"/user/folder/image.png",
		"Image from path",
	)

	_, err = wwaba.Send(ctx, msg)
	if err != nil{
		fmt.Printf("%v",err)
	}
}

```


## Send video message

```go
package main

import (
	"fmt"
	"log"

	wwabago "github.com/Git-RenatoAlcantara/wwabago"
)


func main(){

	wwaba , err := wwabago.CreateWwaba("access_token", "identification_number_id")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	msg := wwabago.NewVideoMessage(
		"+00(00)00000-0000",
		"./video/demo.mp4",
		"Video demo",
	)

	_, err = wwaba.Send(ctx, msg)
	if err != nil{
		fmt.Printf("%v",err)
	}
}

```


