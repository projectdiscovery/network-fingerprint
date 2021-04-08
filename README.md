# network-fingerprint

[![License](https://img.shields.io/badge/license-MIT-_red.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/projectdiscovery/network-fingerprint)](https://goreportcard.com/report/github.com/projectdiscovery/network-fingerprint)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/projectdiscovery/network-fingerprint/issues)
[![GitHub Release](https://img.shields.io/github/release/projectdiscovery/network-fingerprint)](https://github.com/projectdiscovery/naabu/releases)
[![Follow on Twitter](https://img.shields.io/twitter/follow/pdiscoveryio.svg?logo=twitter)](https://twitter.com/pdiscoveryio)
[![Chat on Discord](https://img.shields.io/discord/695645237418131507.svg?logo=discord)](https://discord.gg/KECAGdH)

Capture packet request/response pairs for a port and/or IP to aid in Network protocol based Nuclei Templates creation.

# Resources
- [network-fingerprint](#network-fingerprint)
- [Resources](#resources)
- [Usage](#usage)
- [Installation Instructions](#installation-instructions)
    - [From Binary](#from-binary)
    - [From Source](#from-source)
    - [From Github](#from-github)
- [Running network-fingerprint](#running-network-fingerprint)
    - [Output Format](#output-format)


# Usage

```sh
▶ network-fingerprint -h
```
This will display help for the tool. Here are all the switches it supports.

| Flag           | Description                                            | Example                             |
| -------------- | ------------------------------------------------------ | ----------------------------------- |
| iface              |  Interface to perform capture on (default "lo0")             | network-fingerprint -iface eth0                      |
| ip         |  IP to filter packets for                          | network-fingerprint -ip 127.0.0.1            |
| port              |Port to capture packets on                       | network-fingerprint -port 27017                          |

# Installation Instructions

### From Binary

The installation is easy. You can download the pre-built binaries for your platform from the [releases](https://github.com/projectdiscovery/network-fingerprint/releases/) page. Extract them using tar, move it to your `$PATH`and you're ready to go.

Download latest binary from https://github.com/projectdiscovery/network-fingerprint/releases

```sh
▶ tar -xvf network-fingerprint-linux-amd64.tar
▶ cp naabu-linux-amd64 /usr/local/bin/network-fingerprint
▶ network-fingerprint -h
```

### From Source

naabu requires **go1.14+** to install successfully and have `libpcap-dev` installed on the system.

To install libpcap-dev:-

```sh
apt install -y libpcap-dev
```

```sh
▶ GO111MODULE=on go get -v github.com/projectdiscovery/network-fingerprint
▶ network-fingerprint -h
```

### From Github

```sh
▶ git clone https://github.com/projectdiscovery/network-fingerprint.git; cd network-fingerprint; go build; cp network-fingerprint /usr/local/bin/; network-fingerprint -h
```


# Running network-fingerprint

To run the tool on a target to capture traffic on a port, just use the following command.

```sh
▶ network-fingerprint -port <port>
```

where `<port>` is the port you want to capture traffic for.

To also filter by IP while running on more common ports like 80, where there is a lot of noise, you can use the ip flag.

```sh
▶ network-fingerprint -port <port> -ip <destination-ip> 
```

### Output Format

```
testing@local# ./network-fingerprint -port 27017 -ip 127.0.0.1
2021/04/08 23:15:07 network-packet-capture: nuclei-helper by @pdiscoveryio
2021/04/08 23:15:07 [device] en0 IP: 192.168.1.9
2021/04/08 23:15:07 [device] bridge100 IP: 192.168.64.1
2021/04/08 23:15:07 [device] lo0 IP: 127.0.0.1
```

```json
{
  "data": "\ufffd",
  "hex": "dd",
  "request": true
}
{
  "data": "?\u0001",
  "hex": "3f01",
  "response": true
}
```

Requests (Client to Destination) messages have `request: true` while responses (Destination To Client) have `response: true` set to help in easily identifying correct fingerprints. 