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
	"bufio"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"sync"
)

func OpenConnection(config IrcConfig) (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", config.Server+":"+strconv.Itoa(config.Port))

	if err == nil {
		fmt.Fprintln(conn, "USER "+config.Nick+" 0 * :"+config.RealName)
		fmt.Fprintln(conn, "NICK "+config.Nick)
	}
	return
}

func RecvMsgs(conn net.Conn, wg sync.WaitGroup, config IrcConfig, ircReady chan bool) {
	defer wg.Done()
	sock := bufio.NewReader(conn)
	for {
		msg, err := sock.ReadString('\n')

		if err == nil {
			registered, _ := regexp.MatchString("^:.+ (?:422|376)", msg)
			if registered {
				fmt.Fprintln(conn, "JOIN "+config.Channel)
			}

			joined, _ := regexp.MatchString(".+?End of /NAMES list.", msg)
			if joined {
				ircReady <- true
			}

			fmt.Println(msg)
		}
	}
}

func SendMsg(conn net.Conn, wg sync.WaitGroup, config IrcConfig, channel chan string) {
	defer wg.Done()
	for {
		message := <-channel
		fmt.Fprintf(conn, "PRIVMSG "+config.Channel+" :"+message+"\n")
	}
}
