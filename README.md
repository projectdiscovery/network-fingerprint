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

To run the tool on a target, just use the following command.
```sh
▶ network-fingerprint -host hackerone.com
```

This will run the tool against hackerone.com. There are a number of configuration options that you can pass along with this command. The verbose switch `-v` can be used to display verbose information.

```sh
▶ naabu -host hackerone.com

                  __
  ___  ___  ___ _/ /  __ __
 / _ \/ _ \/ _ \/ _ \/ // /
/_//_/\_,_/\_,_/_.__/\_,_/ v2.0.3

    projectdiscovery.io

[WRN] Use with caution. You are responsible for your actions
[WRN] Developers assume no liability and are not responsible for any misuse or damage.
[INF] Running SYN scan with root privileges
[INF] Found 4 ports on host hackerone.com (104.16.100.52)
hackerone.com:80
hackerone.com:443
hackerone.com:8443
hackerone.com:8080
```
