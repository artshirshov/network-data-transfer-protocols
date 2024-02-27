package main

import (
	"encoding/binary"
	"fmt"
	"github.com/goburrow/modbus"
	"log"
	"time"
)

var (
	tcpAddr       = "109.167.241.225:601"
	salveID       = byte(1)
	studentNumber = 9
	startRegister = uint16(studentNumber * 100)
)

func main() {
	handler := modbus.NewTCPClientHandler(tcpAddr)
	handler.SlaveId = salveID
	defer handler.Close()

	err := handler.Connect()
	if err != nil {
		log.Fatalf("ошибка подключения: %v", err)
	}
	client := modbus.NewClient(handler)

	results, err := client.ReadHoldingRegisters(startRegister, 7)
	if err != nil {
		log.Fatalf("ошибка при чтении данных")
	}

	unsignedNumber := binary.BigEndian.Uint16(results[:2])
	signedNumber := int16(binary.BigEndian.Uint16(results[2:4]))
	asciiString := string(results[4:14])

	fmt.Printf("Беззнаковое число: %d\n", unsignedNumber)
	fmt.Printf("Число со знаком: %d\n", signedNumber)
	fmt.Printf("ASCII строка: %s\n", asciiString)
	fmt.Printf("Текущее время: %v", time.Now())
}
