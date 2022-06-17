# youdu

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

	err := yd.Message().SendText("11111", "test")
	if err != nil {
		log.Fatal(err)
	}
}

```