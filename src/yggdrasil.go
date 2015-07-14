//This file is part of Yggdrasil
//
//Yggdrasil is free software: you can redistribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.
//
//Yggdrasil is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.
//
//You should have received a copy of the GNU General Public License
//along with Yggdrasil.  If not, see <http://www.gnu.org/licenses/>.

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
