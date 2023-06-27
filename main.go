package main

import (
	"fmt"
	"log"
	"wstest/config"

	"github.com/gorilla/websocket"
)
type App struct {
	conn *websocket.Conn
}
func main() {
	app := &App{}
	for j := 0; j < 600; j++ {
		go app.ConnectToWs()

	}

	select {}
}
func (app *App)ConnectToWs() {
	conn, _, err := websocket.DefaultDialer.Dial(config.WsUrl, nil)
	if err != nil {
		log.Fatal("连接失败：", err)
	}
	fmt.Println("连接成功")
	app.conn = conn
	defer conn.Close()
	
}
func (app *App)HeadleData() {
	for {
		_, message, err := app.conn.ReadMessage()
		if err != nil {
			log.Fatal("读取失败：", err)
		}
		fmt.Println("接收到消息", string(message))
	}
}
