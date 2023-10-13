# kratos_wss_demo

一个实例工程，展示如何在 kratos 中对接 websocket + tls (wss)。

## 运行方法

1. 执行 `make` 生成证书。windows 用户可以将 Makefile 中的命令直接复制出来执行。
2. 执行 `go run ./server` 开启服务端。
3. 执行 `go run ./client` 执行客户端。

客户端和服务端都会打印出收到的消息。
