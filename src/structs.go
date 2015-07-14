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
	"strconv"
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
	Type     string    `json:"type"`
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
