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
	"errors"
	"net/http"
	"strconv"
	"sync"
)

func IsHTTPServiceOnline(service Service) (b bool, err error) {
	b = false
	resp, err := http.Get(service.Host + ":" + strconv.Itoa(service.Port))

	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			b = true
		}
	}
	return
}

func CheckServices(config Config, wg sync.WaitGroup, ircReady chan bool, ircOut chan string) {
	defer wg.Done()
	for {
		ready := <-ircReady
		if ready {
			ircOut <- "Initiating full service scan..."
			for i := range config.Services {
				var err error

				service := config.Services[i]

				if service.Type == HTTPService {
					_, err = IsHTTPServiceOnline(service)
				} else {
					_, err = false, errors.New("Unknown service type.")
				}

				if err != nil {
					ircOut <- "Service " + service.Name + " seems to be down: " + err.Error()
				}
			}
		}
	}
}
