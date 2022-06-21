# youdu

[Feature](./todo.md)

## Install

```bash
go get github.com/go-packagist/youdu
```

## Usage

```go
package main

import (
	"github.com/go-packagist/youdu"
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
	yd.Message().Send(&youdu.TextMessage{
		ToUser:  "11111",
		ToDept:  "",
		MsgType: youdu.MsgTypeText,
		Text: &youdu.TextItem{
			Content: "test",
		},
	})

	mediaId, err := yd.Media().Upload(youdu.MediaTypeImage, "test.jpeg")
	if err != nil {
		panic(err)
	}
	yd.Message().Send(&youdu.ImageMessage{
		ToUser:  "11111",
		ToDept:  "",
		MsgType: youdu.MsgTypeImage,
		Image: &youdu.MediaItem{
			MediaId: mediaId,
		},
	})

}

```