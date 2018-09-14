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

import (
	"flag"
	"log"

	"net"
	"istio.io/istio/pkg/adsc"
)

var (
	clients   = flag.Int("clients",
		1,
		"Number of ads clients")

	pilotAddr = flag.String("pilot",
		"localhost:15010",
		//"pilot.v10.istio.webinf.info:15011",
		"Pilot address")

	certDir = flag.String("certDir",
		"", // /etc/certs",
		"Certificate dir")

	verbose = false
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
	adsc, err := adsc.Dial(*pilotAddr, *certDir, &adsc.ADSCOpt{
		IP: net.IPv4(10, 10, byte(n / 256), byte(n % 256)).String(),
	})
	if err != nil {
		log.Println("Error connecting ", err)
		return
	}

	//t0 := time.Now()

	adsc.Watch()

	//adsc.SendRsc(v2.RouteType, "ingress~10.10.10.10~istio-ingress-794f555d7b-5fm48.istio-system~istio-system.svc.cluster.local-36", []string{
	//	"http.80",
	//	"http.443",
	//})

	//for {
		//msg, err := adsc.Recv(15 * time.Second)
		//if err != nil {
		//	log.Println("Stream closed:", err)
		//	return
		//} else {
		//	for _, rsc := range msg.Resources {
		//		if verbose {
		//			log.Println(msg.VersionInfo, rsc.TypeUrl)
		//			if rsc.TypeUrl == v2.ListenerType {
		//				valBytes := rsc.Value
		//				ll := &xdsapi.Listener{}
		//				proto.Unmarshal(valBytes, ll)
		//
		//				tm := &jsonpb.Marshaler{Indent: "  "}
		//				log.Println(tm.MarshalToString(ll))
		//			} else if rsc.TypeUrl == v2.RouteType {
		//				valBytes := rsc.Value
		//				ll := &xdsapi.RouteConfiguration{}
		//				proto.Unmarshal(valBytes, ll)
		//
		//				tm := &jsonpb.Marshaler{Indent: "  "}
		//				log.Println(tm.MarshalToString(ll))
		//			}
		//		}
		//	}
		//	log.Println("Received ", len(msg.Resources), time.Since(t0))
		//	adsc.Ack(msg)
		//}
	}

