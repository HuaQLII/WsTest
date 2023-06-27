package config

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func main1() {
	for i := 0; i < 1000; i++ {
		go ConnectToWs()

	}
	select {}
}
func ConnectToWs() {
	conn, _, err := websocket.DefaultDialer.Dial(WsUrl, nil)
	if err != nil {
		log.Fatal("连接失败：", err)
	}
	fmt.Println("连接成功")
	defer conn.Close()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatal("读取失败：", err)
		}
		fmt.Println("接收到消息", string(message))
	}
}
