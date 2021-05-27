// Copyright Mukul Agarwal 2021

// This program is free software: you can redistribute it and/or modify
//     it under the terms of the GNU Affero General Public License as
//     published by the Free Software Foundation, either version 3 of
//     the License, or (at your option) any later version.
//
//     This program is distributed in the hope that it will be useful,
//     but WITHOUT ANY WARRANTY; without even the implied warranty of
//     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//     GNU General Public License for more details.
//
//     You should have received a copy of the GNU Affero General Public License
//     along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/fatih/color"
)

var (
	port   int  = 8080
	header bool = false
)

func loadPortData() string {
	rawPortData := os.Getenv("PORT")
	if rawPortData != "" {
		parsedPort, err := strconv.Atoi(rawPortData)
		if err == nil {
			if parsedPort >= 0 && parsedPort <= 65535 {
				port = parsedPort
			} else {
				color.Yellow(`Given port %d not within range 0...65535
				Defaulting to port %d`, parsedPort, port)
			}
		} else {
			color.Yellow(`Error parsing port data (%v): %v
			Defaulting to port %d`, rawPortData, err, port)
		}
	}
	return fmt.Sprintf("0.0.0.0:%d", port)
}

func loadHeaderData() {
	rawHeaderData := os.Getenv("HEADER")
	if rawHeaderData != "" {
		header = true
	}
}

func handle_request(w http.ResponseWriter, req *http.Request) {
	color.Green("%s %s", req.Method, req.URL.Path)
	if header {
		for name, headers := range req.Header {
			for _, h := range headers {
				color.Cyan("\t%v: %v", name, h)
			}
		}
	}
	w.WriteHeader(404) // not found
}

func main() {
	portString := loadPortData()
	http.HandleFunc("/", handle_request)

	go color.Green("Listening at port %d", port)
	err := http.ListenAndServe(portString, nil)
	if err != nil {
		color.Red("Error listening and serving: %v", err)
	}
}
