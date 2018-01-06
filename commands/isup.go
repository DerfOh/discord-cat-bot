package command

import (
	"fmt"
	"net"
	"time"
)

//IsUp returns the status of a host
func IsUp(content []string) string {
	host := content[1]
	seconds := 5
	timeOut := time.Duration(seconds) * time.Second

	conn, err := net.DialTimeout("tcp", host, timeOut)

	if err != nil {
		fmt.Println(err)
		return "Unable to connect to " + host + " " + err.Error()
	}

	response := "Connection established to " + host + " (" + conn.RemoteAddr().String() + ")"
	return response
}
