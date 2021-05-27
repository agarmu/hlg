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

func main() {
	var port int = 8080
	var header bool = false

	rawPortData := os.Getenv("PORT")
	rawHeaderData := os.Getenv("HEADER")
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
	portString := fmt.Sprintf(":%d", port)

	if rawHeaderData != "" {
		header = true
	}

	handle_request := func(w http.ResponseWriter, req *http.Request) {
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

	http.HandleFunc("/", handle_request)

	color.Green("Listening at port %d", port)
	fmt.Println()
	http.ListenAndServe(portString, nil)
	defer func() {
		fmt.Println()
		color.HiWhite(`hlg - The Http Logger
		Copyright 2021 Mukul Agarwal
		Released under the GNU Affero General Public License, available at <https://www.gnu.org/licenses/>.`)
	}()
}
