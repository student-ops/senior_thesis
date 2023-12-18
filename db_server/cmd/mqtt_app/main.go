package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	fmt.Println("*** 開始 ***")

	// バッファサイズを持つチャネルの作成
	msgCh := make(chan mqtt.Message, 100)
	listeningTopic := "test"
	brokerAddress := os.Getenv("MQTT_BROKER_ADDRESS")

	// メッセージハンドラの定義
	var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		go func(m mqtt.Message) {
			// バッファ付きチャネルにメッセージを送信
			msgCh <- m
		}(msg)
	}

	// MQTTクライアントの設定
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerAddress)
	opts.SetDefaultPublishHandler(f)
	cc := mqtt.NewClient(opts)

	// MQTTブローカーへの接続
	if token := cc.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Mqtt error: %s", token.Error())
	}

	// トピックへのサブスクライブ
	if subscribeToken := cc.Subscribe(listeningTopic, 0, f); subscribeToken.Wait() && subscribeToken.Error() != nil {
		log.Fatal(subscribeToken.Error())
	}

	// OSのシグナルハンドリング
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	// メッセージの受信と処理
	for {
		select {
		case m := <-msgCh:
			// チャネルからメッセージを取り出し、処理
			fmt.Printf("topic: %v, payload: %v\n", m.Topic(), string(m.Payload()))
		case <-signalCh:
			// 中断シグナルの検出
			fmt.Println("Interrupt detected.")
			cc.Disconnect(1000)
			return
		}
	}
}
