// Copyright (c) 2016 struktur AG.
// Use of this source code is governed by the MIT License that can be
// that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	listenIP := flag.String("ip", "", "Listen IP for the DNS server, also for A and PTR request")
	listenPort := flag.Int("port", 53, "Listen port for the DNS server")
	name := flag.String("name", "", "A host record")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <name>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if *listenIP == "" || *name == "" {
		flag.Usage()
		os.Exit(1)
	}

	server, err := NewServer(*listenIP, *listenPort)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = server.Serve(*name)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print("exiting")
}
