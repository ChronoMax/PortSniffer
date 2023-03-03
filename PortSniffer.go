package main

import (
	"fmt"

	//Ref to the port directory
	"github.com/ChronoMax/PortSniffer/port"
)

func main() {
	fmt.Println("Port Scanner..")

	//Difining the var open as the return value of the func ScanPort in the port directory.
	open := port.ScanPort("tcp", "localhost", 80)
	//Prints the text with the set value.
	fmt.Printf("Port open: %t\n", open)
}
