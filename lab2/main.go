package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

var (
	// Connection config
	broker = "broker.hivemq.com"
	port   = 1883

	// Messages
	myMessages  []messageMeta
	allMessages []messageMeta
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	routeMask := strings.Split(msg.Topic(), "/")

	if routeMask[1] == "STUDENT21" &&
		(routeMask[2] == "Value1" || routeMask[2] == "Value2") {
		fmt.Printf("My topics ITMO/Student21/%s  \tReceived message: %s\n", routeMask[2], msg.Payload())
		myMessages = append(myMessages, create(msg))
	} else if routeMask[2] == "Value3" {
		fmt.Printf("All topics ITMO/%s/Value3 \tReceived message: %s\n", routeMask[1], msg.Payload())
		allMessages = append(allMessages, create(msg))
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main() {
	opts := getClientOpt()
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	go sub(client, "ITMO/#")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	client.Disconnect(250)

	fmt.Println("\nMy messages:")
	for _, message := range myMessages {
		fmt.Println(message)
	}

	fmt.Println("\nAll messages:")
	for _, message := range allMessages {
		fmt.Println(message)
	}
}

func getClientOpt() *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	return opts
}

func sub(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
}

type messageMeta struct {
	route     string
	studentId int
	valueId   int
	msg       string
}

func create(msg mqtt.Message) messageMeta {
	routeMask := strings.Split(msg.Topic(), "/")
	studentId, _ := strconv.Atoi(routeMask[1][7:])
	valueId, _ := strconv.Atoi(routeMask[2][5:])

	return messageMeta{
		route:     msg.Topic(),
		studentId: studentId,
		valueId:   valueId,
		msg:       string(msg.Payload()),
	}
}

func (mm messageMeta) String() string {
	return fmt.Sprintf(
		"Route: %s,\tStudent ID: %d\t Value ID: %d\tValue: %s",
		mm.route,
		mm.studentId,
		mm.valueId,
		mm.msg,
	)
}
