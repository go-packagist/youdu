# youdu

<div align="center" style="padding: 30px 0;">
    <img src="logo.png" width="240">
    <p>Golang Youdu SDK</p>
</div>


[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/youdu)](https://goreportcard.com/report/github.com/go-packagist/youdu)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/youdu)](https://pkg.go.dev/github.com/go-packagist/youdu)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)


[Feature](./todo.md)

## Installation

```bash
go get github.com/go-packagist/youdu
```

## Usage

```go
package main

import (
	"github.com/go-packagist/youdu"
	"github.com/go-packagist/youdu/message"
	"log"
)

func main() {
	yd := youdu.New(&youdu.Config{
		Api:    "http://domain.com/api",
		Buin:   1111111,
		AppId:  "22222222222222",
		AesKey: "3444444444444444444444444444444444",
	})

	yd.Message().SendText("11111", "test")
	yd.Message().Send(&message.TextMessage{
		ToUser:  "11111",
		ToDept:  "",
		MsgType: message.MsgTypeText,
		Text: &message.TextItem{
			Content: "test",
		},
	})

	mediaId, err := yd.Media().Upload(youdu.MediaTypeImage, "test.jpeg")
	if err != nil {
		panic(err)
	}
	yd.Message().Send(&message.ImageMessage{
		ToUser:  "11111",
		ToDept:  "",
		MsgType: message.MsgTypeImage,
		Image: &message.MediaItem{
			MediaId: mediaId,
		},
	})

}

```