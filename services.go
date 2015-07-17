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
	"net/smtp"
	"strconv"
	"sync"
)

func GetServiceAddress(host string, port int) (address string) {
	address = host + ":" + strconv.Itoa(port)
	return
}

func IsHTTPServiceOnline(service Service) (b bool, err error) {
	b = false
	resp, err := http.Get(GetServiceAddress(service.Host, service.Port))

	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			b = true
		}
	}
	return
}

func IsFTPServiceOnline(service Service) (b bool, err error) {
	b = false
	err = nil
	return
}

func IsSSHServiceOnline(service Service) (b bool, err error) {
	b = false
	err = nil
	return
}

func IsSMTPServiceOnline(service Service) (b bool, err error) {
	b = false
	client, err := smtp.Dial(GetServiceAddress(service.Host, service.Port))

	if err == nil {
		defer client.Close()
		b = true
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
				var online bool

				service := config.Services[i]

				if service.Type == HTTPService {
					online, err = IsHTTPServiceOnline(service)
				} else if service.Type == FTPService {
					online, err = IsFTPServiceOnline(service)
				} else if service.Type == SSHService {
					online, err = IsSMTPServiceOnline(service)
				} else if service.Type == SMTPService {
					online, err = IsSMTPServiceOnline(service)
				} else {
					online, err = false, errors.New("Unknown service type.")
				}

				if online {
					// Log response time to sqlite db
				}
				if err != nil {
					ircOut <- "Service " + service.Name + " seems to be down: " + err.Error()
				}
			}
			ircOut <- "Scan done."
		}
	}
}
