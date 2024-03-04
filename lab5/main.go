package main

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
)

var (
	ipAddr = "109.167.241.225"
	port   = 61556
)

func main() {
	sendMessage()
	buf := getMessage()
	parseMessage(buf)
}

func sendMessage() {
	con, err := net.Dial("udp", fmt.Sprintf("%s:%d", ipAddr, port+1))
	if err != nil {
		log.Fatalln(err)
	}
	defer con.Close()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			msg := strconv.Itoa(n) + " Shirshov student 21"
			_, err := con.Write([]byte(msg))
			if err != nil {
				log.Fatalln(err)
			}
		}(i)
	}
	wg.Wait()
}

func getMessage() []byte {
	cl, err := ftp.Dial(ipAddr + ":21")
	if err != nil {
		log.Fatalln(err)
	}
	err = cl.Login("Student", "FksG5$%^rgtdSDFH")
	if err != nil {
		log.Fatalln(err)
	}
	r, err := cl.Retr("UDP log.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Close()

	buf, err := io.ReadAll(r)
	if err != nil {
		log.Fatalln(err)
	}
	err = cl.Quit()
	if err != nil {
		log.Fatalln(err)
	}
	return buf
}

func parseMessage(buf []byte) {
	msg := string(buf)

	raws := strings.Split(msg, "\n")

	for _, raw := range raws {
		if strings.Contains(raw, "Shirshov student 21") {
			fmt.Println(raw)
		}
	}
}
