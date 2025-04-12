package capture

import (
	"dns-tracker/model"
	"dns-tracker/writer"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func ListenDNS(interfaceName string, w *writer.JSONWriter) error {
	handle, err := pcap.OpenLive(interfaceName, 65535, true, pcap.BlockForever)
	if err != nil {
		return err
	}
	defer handle.Close()

	var filter = "udp and port 53"
	if err = handle.SetBPFFilter(filter); err != nil {
		return err
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		dnsLayer := packet.Layer(layers.LayerTypeDNS)
		if dnsLayer == nil {
			continue
		}

		dns, _ := dnsLayer.(*layers.DNS)
		if dns.OpCode != layers.DNSOpCodeQuery || len(dns.Questions) == 0 {
			continue
		}

		ipLayer := packet.NetworkLayer()
		if ipLayer == nil {
			continue
		}

		srcIP := ipLayer.NetworkFlow().Src().String()
		q := dns.Questions[0]

		logEntry := model.DNSLog{
			Timestamp: time.Now(),
			SrcIP:     srcIP,
			Query:     string(q.Name),
			QType:     q.Type.String(),
		}

		if err = w.Write(logEntry); err != nil {
			log.Println("write error:", err)
		}
	}
	return nil
}
