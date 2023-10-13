package main

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"net/url"

	"nhooyr.io/websocket"
)

func main() {
	// 忽略证书错误。正式服务上不用加。
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	ctx := context.Background()
	u := &url.URL{
		Scheme: "wss",
		Host:   "localhost:8080",
		Path:   "/wss",
	}

	c, _, err := websocket.Dial(ctx, u.String(), &websocket.DialOptions{
		HTTPClient: client,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close(websocket.StatusNormalClosure, "")

	err = c.Write(ctx, websocket.MessageText, []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	_, message, err := c.Read(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(message))
}
