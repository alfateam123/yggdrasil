package main

import (
	"fmt"
	"net"
	"strconv"
)

type Service struct {
	Host    string
	Port    int
	Address string
	Proto   string
	Name    string
}

func NewService(host string, port int, proto string, name string) (service Service) {
	service = Service{
		Host:    host,
		Port:    port,
		Address: host + ":" + strconv.Itoa(port),
		Proto:   proto,
		Name:    name,
	}
	return service
}

func main() {
	var services []Service
	services = append(services, NewService("nwa.xyz", 80, "tcp", "NWA Web Server"))
	services = append(services, NewService("nerdz.eu", 80, "tcp", "NERDZ Web Server"))
	services = append(services, NewService("google.com", 80, "tcp", "GOOGLE Web Server"))
	services = append(services, NewService("wikipedia.org", 80, "tcp", "WIKIPEDIA Web Server"))
	services = append(services, NewService("microsoft.com", 80, "tcp", "MICROSOFT Web Server"))

	for i := range services {
		fmt.Println(services[i].Proto + " " + services[i].Address)
		conn, err := net.Dial(services[i].Proto, services[i].Address)
		if err != nil {
			panic(err)
		} else {
			conn.Close()
			fmt.Println("Tested " + services[i].Name)
		}
	}
}
