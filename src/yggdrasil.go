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
	"io"
	"os"
	"sync"
	"time"
)

func GetConfig(r io.Reader) (x *Config, err error) {
	x = new(Config)
	err = json.NewDecoder(r).Decode(x)
	return
}

func Timer(interval int, ircReady chan bool, timeout chan bool) {
	ready := <-ircReady
	if ready {
		for {
			time.Sleep(time.Duration(interval) * time.Second)
			timeout <- true
		}
	}
}

func main() {
	var wg sync.WaitGroup
	conffile, err := os.Open("./config.json")
	if err != nil {
		panic(err)
	}

	config, err := GetConfig(conffile)
	if err != nil {
		panic(err)
	}

	conn, err := OpenConnection(config.Irc)
	if err == nil {
		ircOut := make(chan string)
		ircReady := make(chan bool, 1)
		timeout := make(chan bool)
		wg.Add(3)
		go RecvMsgs(conn, wg, config.Irc, ircReady)
		go SendMsg(conn, wg, config.Irc, ircOut)
		go CheckServices(*config, wg, timeout, ircOut)
		go Timer(config.Interval, ircReady, timeout)
	}
	wg.Wait()
}
