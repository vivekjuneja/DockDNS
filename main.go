// DockDNS, the simple docker-aware DNS forwarder.
// Copyright 2014 Vladimir "farcaller" Pouzanov <farcaller@gmail.com>
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
	"github.com/farcaller/dockdns/resolver"
	"github.com/farcaller/dockdns/server"
	"github.com/fsouza/go-dockerclient"
	"log"
	"os"
	"os/signal"
	"fmt"
)

var forwardDNS = flag.String("forward", "8.8.8.8:53", "IP address of forwarder DNS")
var dockerZone = flag.String("zone", "docker.", "Docker zone name")
var dockerEndpoint = flag.String("docker", "https://172.17.42.1:2376", "Docker API endpoint")
var listenAddress = flag.String("listen", "172.17.0.4:53", "DNS listen address")
var certPath = flag.String("cert", "./cert.pem", "Path to the Certificate PEM file")
var keyPath = flag.String("key", "./key.pem", "Path to the Key PEM file")

func main() {
	flag.Parse()

	dockerClient, error := docker.NewTLSClient(*dockerEndpoint, *certPath, *keyPath, "")
	
	if error != nil {	
		fmt.Println(error)
	}

	dockerResolver := resolver.NewDocker(dockerClient, *dockerZone)
	resolver := resolver.NewForward(*forwardDNS)
	server := server.New(*dockerZone, *listenAddress, dockerResolver, resolver)

	server.Run()
	log.Printf("Server listening on TCP/UDP %s\n", *listenAddress)

  	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)


	for {
		select {
		case <-sig:
			return
		}
	}
	
}
