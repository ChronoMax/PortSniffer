package main

import (
	"fmt"

	"github.com/ChronoMax/PortSniffer/port"
)

func main() {
	fmt.Println("Port Scanner..")

	open := port.ScanPort("tcp", "localhost", 80)
	fmt.Printf("Port open: %t\n", open)
}
