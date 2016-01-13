xudnsd
==========

A very simple single host/IP DNS server.

- Serves a single A host with the IP address it listens on.
- Reverse lookup of the IP address returns the A host.

## Running

```bash
xudnsd -name=myhost.local. -ip=192.168.1.54 -port=1053
```

## Test

```bash
nslookup -port=1053 -debug myhost.local 192.168.1.54
nslookup -port=1053 -debug 192.168.1.54 192.168.1.54
```

## License
xudnsd is governed by the MIT License. See the LICENSE file for details.
