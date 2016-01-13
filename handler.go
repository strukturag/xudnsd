// Copyright (c) 2016 struktur AG.
// Use of this source code is governed by the MIT License that can be
// that can be found in the LICENSE file.

package main

import (
	"github.com/miekg/dns"
	"net"
)

// A Handler represents the DNS handler with IP, hostname.
type Handler struct {
	name    string
	ip      net.IP
	reverse string
}

// NewHandler returns a Handler filled from name and ip.
func NewHandler(name, ip string) *Handler {
	reverse, _ := dns.ReverseAddr(ip)
	return &Handler{
		name:    name,
		ip:      net.ParseIP(ip),
		reverse: reverse,
	}
}

func (h *Handler) handleQuery(w dns.ResponseWriter, r *dns.Msg) {
	q := r.Question[0]

	m := &dns.Msg{}
	m.SetReply(r)

	switch q.Qtype {
	case dns.TypePTR:
		if q.Name == h.reverse {
			rr := &dns.PTR{}
			rr.Hdr = dns.RR_Header{Name: q.Name, Rrtype: dns.TypePTR, Class: dns.ClassINET, Ttl: 0}
			rr.Ptr = h.name
			m.Answer = append(m.Answer, rr)
		}
	case dns.TypeA:
		if q.Name == h.name {
			rr := &dns.A{}
			rr.Hdr = dns.RR_Header{Name: q.Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 0}
			rr.A = h.ip
			m.Answer = append(m.Answer, rr)
		}
	}

	w.WriteMsg(m)
}
