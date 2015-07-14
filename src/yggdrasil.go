package main

import (
	"fmt"
	"net"
	"strconv"
	"encoding/json"
	"os"
	"io"
)

type Config struct {
	Services []Service `json:"services"`
	Irc      IrcConfig `json:"irc"`
}

type IrcConfig struct {
	Nick     string    `json:"nick"`
	Host     string    `json:"host"`
	Port     int       `json:"port"`
	Chan     string    `json:"chan"`
}

type Service struct {
	Host     string    `json:"host"`
	Port     int       `json:"port"`
	Address  string
	Proto    string    `json:"proto"`
	Name     string    `json:"name"`
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

func GetConfig(r io.Reader) (x *Config, err error) {
	x = new(Config)
	err = json.NewDecoder(r).Decode(x)
	return
}

func main() {
	conffile, err := os.Open("./config.json")
	if err != nil {
		panic(err)
	}

	config, err := GetConfig(conffile)
	if err != nil {
		panic(err)
	}

	for i := range config.Services {
		service := config.Services[i]
		service.Address = service.Host + ":" + strconv.Itoa(service.Port)
		fmt.Println(service.Proto + " " + service.Address)
		conn, err := net.Dial(service.Proto, service.Address)
		if err != nil {
			panic(err)
		} else {
			conn.Close()
			fmt.Println("Tested " + service.Name)
		}
	}
}
