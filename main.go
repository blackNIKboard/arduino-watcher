package main

import (
	"flag"
	"fmt"
	serial "github.com/tarm/serial"
	"log"
	"strconv"
	"time"
)

var device = flag.String("device", "", "device to call")

func main() {
	flag.Parse()

	c := &serial.Config{
		Name: *device,
		Baud: 9600,
		//		ReadTimeout: time.Millisecond * 5,
	}

	port, err := serial.OpenPort(c)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 100; i++ {
		n, err := port.Write([]byte("test" + strconv.Itoa(i) + "\n"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Sent %v bytes\n", n)

		time.Sleep(time.Millisecond * 500)
	}
	time.Sleep(time.Second * 3)
	port.Close()
}

