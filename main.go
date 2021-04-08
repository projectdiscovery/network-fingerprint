package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/tidwall/pretty"
)

var (
	iface = flag.String("iface", "lo0", "Interface to perform capture on")
	port  = flag.String("port", "", "Port to capture packets on")
	ip    = flag.String("ip", "", "IP to filter packets for")
)

var buffer = int32(65535)

type capture struct {
	Data     string `json:"data"`
	Hex      string `json:"hex"`
	Request  bool   `json:"request,omitempty"`
	Response bool   `json:"response,omitempty"`
}

func main() {
	flag.Parse()

	log.Printf("network-fingerprint: nuclei-helper by @pdiscoveryio")
	if *port == "" {
		log.Fatalf("No port provided! Exiting.")
	}

	var bpfFilter string
	if *ip == "" {
		bpfFilter = "tcp and port " + *port
	} else {
		bpfFilter = "tcp and host " + *ip + " and port " + *port
	}

	found := deviceExists(*iface)
	if !found {
		log.Fatalf("No capture device found for iface. Exiting!")
	}

	handler, err := pcap.OpenLive(*iface, buffer, false, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handler.Close()

	if err := handler.SetBPFFilter(bpfFilter); err != nil {
		log.Fatal(err)
	}

	source := gopacket.NewPacketSource(handler, handler.LinkType())
	for packet := range source.Packets() {
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if tcpLayer == nil {
			continue
		}
		tcp, _ := tcpLayer.(*layers.TCP)

		app := packet.ApplicationLayer()
		if app != nil {
			payload := app.Payload()

			cap := &capture{
				Data: string(payload),
				Hex:  hex.EncodeToString(payload),
			}
			if strconv.Itoa(int(tcp.SrcPort)) == *port {
				cap.Response = true
			} else {
				cap.Request = true
			}
			data, err := json.Marshal(cap)
			if err != nil {
				log.Printf("[error] could not marshal request: %s\n", err)
			}
			data = pretty.Color(pretty.Pretty(data), nil)
			os.Stdout.Write(data)
		}
	}
}

// deviceExists lists all devices and also returns true if the specified
// device with the name exists.
func deviceExists(name string) bool {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalf("Could not find devices: %v\n", err)
	}
	var matched bool
	for _, device := range devices {
		if device.Name == name {
			matched = true
		}

		var ips []string
		for _, ip := range device.Addresses {
			ip4 := ip.IP.To4().String()
			if ip4 != "" && ip4 != "<nil>" {
				ips = append(ips, ip4)
			}
		}
		if len(ips) == 0 {
			continue
		}
		log.Printf("[device] %s IP: %v\n", device.Name, strings.Join(ips, ","))
	}
	return matched
}
