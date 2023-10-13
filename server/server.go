package main

import (
	"crypto/tls"
	"log"
	nt "net/http"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"nhooyr.io/websocket"
)

func main() {
	conf, err := makeConf("tls.crt", "tls.key")
	if err != nil {
		log.Fatal(err)
	}

	httpSrv := http.NewServer(
		http.Address(":8080"),
		http.TLSConfig(conf),
	)
	httpSrv.HandleFunc("/wss", wss) // <- 其实普通 web 工程咋对接 kratos 就咋对接

	app := kratos.New(
		kratos.Name("ws"),
		kratos.Server(
			httpSrv,
		),
	)
	if err := app.Run(); err != nil {
		log.Println(err)
	}
}

func makeConf(certFile string, keyFile string) (*tls.Config, error) {
	var conf tls.Config
	conf.InsecureSkipVerify = true // <- 正式服务上不用加

	c, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	conf.Certificates = []tls.Certificate{c}
	return &conf, nil
}

func wss(w nt.ResponseWriter, r *nt.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer c.Close(websocket.StatusNormalClosure, "")

	ctx := r.Context()
	for {
		mt, message, err := c.Read(ctx)
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)

		err = c.Write(ctx, mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
