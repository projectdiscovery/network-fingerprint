# network-fingerprint

[![License](https://img.shields.io/badge/license-MIT-_red.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/projectdiscovery/network-fingerprint)](https://goreportcard.com/report/github.com/projectdiscovery/network-fingerprint)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/projectdiscovery/network-fingerprint/issues)
[![Follow on Twitter](https://img.shields.io/twitter/follow/pdnuclei.svg?logo=twitter)](https://twitter.com/pdnuclei)
[![Chat on Discord](https://img.shields.io/discord/695645237418131507.svg?logo=discord)](https://discord.gg/projectdiscovery)

Capture packet request/response pairs for a port and/or IP to aid in Network protocol based Nuclei Templates creation.

# Resources
- [network-fingerprint](#network-fingerprint)
- [Resources](#resources)
- [Usage](#usage)
- [Installation Instructions](#installation-instructions)
  - [To install `libpcap-dev`:](#to-install-libpcap-dev)
    - [On Debian and its derivatives](#on-debian-and-its-derivatives)
    - [On Arch Based Systems](#on-arch-based-systems)
    - [On Mac install using homebrew](#on-mac-install-using-homebrew)
    - [On Windows](#on-windows)
  - [Installing Network-fingerprint](#installing-network-fingerprint)
- [Running network-fingerprint](#running-network-fingerprint)
    - [Output Format](#output-format)


# Usage

```sh
▶ network-fingerprint -h
```
This will display help for the tool. You can find all the supported switches below:

| Flag  | Description                                     | Example                             |
|-------|-------------------------------------------------|-------------------------------------|
| iface | Interface to perform capture on (default `lo0`) | `network-fingerprint -iface eth0`   |
| ip    | IP to filter packets for                        | `network-fingerprint -ip 127.0.0.1` |
| port  | Port to capture packets on                      | `network-fingerprint -port 27017`   |


# Installation Instructions


`network-fingerprint` requires **go1.18+** to install successfully and have `libpcap-dev` installed on the system.

## To install `libpcap-dev`:

### On Debian and its derivatives

```sh
▶ sudo apt install -y libpcap-dev
```

### On Arch Based Systems

```sh
▶ sudo pacman -S libpcap
```

### On Mac install using homebrew

```sh
▶ brew install libpcap
```

### On Windows 

Install Npcap from [here](https://npcap.com)


## Installing Network-fingerprint

- Download from Releases

- Install from Source

```sh
▶ go install -v github.com/projectdiscovery/network-fingerprint@latest
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

Here is a detailed blog showcasing the uses of network-fingerprint - https://blog.projectdiscovery.io/writing-network-templates-with-nuclei/

### Output Format

```bash
testing@local# network-fingerprint -port 27017 -ip 127.0.0.1
2021/04/08 23:15:07 network-fingerprint: nuclei-helper by @pdiscoveryio
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
