package port

import (
	"net"
	"strconv"
	"time"
)

// Creating a struct
type ScanResult struct {
	Port  int
	State string
}

func ScanPort(protocol, hostname string, port int) ScanResult {
	//Initialize struct
	result := ScanResult{Port: port}

	//Create adress hostname with port.
	address := hostname + ":" + strconv.Itoa(port)
	//if there is a timeout after 60 seconds.
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	//if the connection does not accept hostname with port retun false.
	if err != nil {
		//Sets the state of struct as following.
		result.State = "Closed"
		return result
	}
	//if connection does accept hostname with port, close connection and retun true.
	defer conn.Close()

	//Sets the state of struct as following.
	result.State = "Open"
	return result
}

func InitialScan(hostname string) []ScanResult {
	//create array based on scanresult array.
	var result []ScanResult

	//Loops for every port till port 1024.
	for i := 1; i < 1024; i++ {
		//Sets the result.	(protocol, hostname, ip)
		result = append(result, ScanPort("tcp", hostname, i))
	}
	return result
}
