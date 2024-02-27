package main

import (
	"fmt"
	"log"
	"math"
	"net"
	"regexp"
	"time"
)

var (
	ipAddr = "109.167.241.225"
	port   = 6340
)

func main() {
	con, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ipAddr, port))
	if err != nil {
		log.Fatalln(err)
	}
	defer con.Close()

	res := make([]byte, 0)
	buf := make([]byte, math.MaxUint16+1)

	for {
		n, err := con.Read(buf)
		if err != nil {
			log.Fatalln(err)
		}

		if n != 4 {
			res = append(res, buf...)
			break
		}
		time.Sleep(time.Second * 1)
	}

	matched, _ := regexp.Compile("Student21\\s\\S+\\s\\S+")
	results := matched.FindAllString(string(res), 10)

	for _, val := range results {
		fmt.Println(val)
	}
}
