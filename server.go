// Copyright (c) 2016 struktur AG.
// Use of this source code is governed by the MIT License that can be
// that can be found in the LICENSE file.

package main

import (
	"github.com/miekg/dns"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

// A Server represents a DNS server.
type Server struct {
	ip   string
	port int
}

// NewServer returns a Server created with IP and port.
func NewServer(ip string, port int) (*Server, error) {
	return &Server{ip, port}, nil
}

// Addr returns the listen address (IP:port) as string.
func (s *Server) Addr() string {
	return s.ip + ":" + strconv.Itoa(s.port)
}

// Serve starts the TCP and UDP listeners and blocks until
// syscall.SIGINT or syscall.SIGTERM is received.
func (s *Server) Serve(name string) (err error) {
	log.Printf("creating DNS service for %s", name)
	handler := NewHandler(name, s.ip)
	dns.HandleFunc(".", handler.handleQuery)
	tcpServer := &dns.Server{
		Addr: s.Addr(),
		Net:  "tcp",
	}
	udpServer := &dns.Server{
		Addr:    s.Addr(),
		Net:     "udp",
		UDPSize: 65535,
	}
	go s.start(udpServer)
	go s.start(tcpServer)

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	return
}

func (s *Server) start(ds *dns.Server) {
	log.Printf("start %s listener on %s\n", ds.Net, s.Addr())
	err := ds.ListenAndServe()
	if err != nil {
		log.Printf("start %s listener on %s failed:%s\n", ds.Net, s.Addr(), err.Error())
	}
}
