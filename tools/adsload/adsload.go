// Copyright 2018 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

// Will create the specified number of ADS connections to the pilot, and keep
// receiving notifications. This creates a load on pilot - without having to
// run a large number of pods in k8s.
// The clients will use 10.10.x.x addresses - it is possible to create ServiceEntry
// objects so clients get inbound listeners. Otherwise only outbound config will
// be pushed.

import (
	"flag"
	"log"
	"time"

	"net"

	"istio.io/istio/pkg/adsc"
)

var (
	clients = flag.Int("clients",
		100,
		"Number of ads clients")

	pilotAddr = flag.String("pilot",
		"localhost:15010",
		"Pilot address. Can be a real pilot exposed for mesh expansion.")

	certDir = flag.String("certDir",
		"", // /etc/certs",
		"Certificate dir. Must be set according to mesh expansion docs for testing a meshex pilot.")
)

func main() {
	flag.Parse()

	for i := 0; i < *clients; i++ {
		n := i
		go runClient(n)
	}
	select {}
}

// runClient creates a single long lived connection
func runClient(n int) {
	c, err := adsc.Dial(*pilotAddr, *certDir, &adsc.Config{
		IP: net.IPv4(10, 10, byte(n/256), byte(n%256)).String(),
	})
	if err != nil {
		log.Println("Error connecting ", err)
		return
	}

	t0 := time.Now()

	c.Watch()

	_, err = c.Wait("rds", 30*time.Second)
	if err != nil {
		log.Println("Timeout receiving RDS")
	}

	log.Println("Initial connection: ", time.Since(t0))

	for {
		msg, err := c.Wait("", 15*time.Second)
		if err == adsc.TimeoutError {
			continue
		}
		log.Println("Received ", msg)
	}
}
