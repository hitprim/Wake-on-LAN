package wol

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"net"
	"strings"
)

func buildMagicPacket(mac string) ([]byte, error) {
	smac := strings.ReplaceAll(mac, ":", "")
	smac = strings.ReplaceAll(smac, "-", "")
	byte_mac, err := hex.DecodeString(smac)
	if err != nil {
		return nil, err
	}
	if len(byte_mac) != 6 {
		return nil, fmt.Errorf("неверный MAC адрес: %s", mac)
	}

	packet := bytes.Repeat([]byte{0xFF}, 6)

	for i := 0; i < 16; i++ {
		packet = append(packet, byte_mac...)
	}

	return packet, nil
}

func Send(mac string, broadcast string, port int) error {
	MAC, err := buildMagicPacket(mac)
	if err != nil {
		return err
	}
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", broadcast, port))
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(MAC)
	if err != nil {
		return err
	}

	return nil
}
