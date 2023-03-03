package port

import (
	"net"
	"strconv"
	"time"
)

func ScanPort(protocol, hostname string, port int) bool {
	//Create adress hostname with port.
	addres := hostname + ":" + strconv.Itoa(port)
	//if there is a timeout after 60 seconds.
	conn, err := net.DailTimeout(protocol, addres, 60*time.Seccond)

	//if the connection does not accept hostname with port retun false.
	if err != nil {
		return false
	}
	//if connection does accept hostname with port, close connection and retun true.
	defer conn.Close()

	return true
}
