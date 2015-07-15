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
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

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
		var err error
		var online bool

		service := config.Services[i]
		service.Address = service.Host + ":" + strconv.Itoa(service.Port)

		if service.Type == "http" {
			online, err = IsHTTPServiceOnline(service)
		} else {
			online, err = false, nil
		}

		if err != nil {
			fmt.Println("Got an error while testing " + service.Name + ": " + err.Error())
		} else if online {
			fmt.Println("Tested " + service.Name)
		}
	}
}
