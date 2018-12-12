package broadlink

import (
	"net"

	"github.com/kihamo/boggart/components/boggart/providers/broadlink/internal"
)

type RemoteType int

const (
	KindRM3Mini   = 0x2737
	KindRMProPlus = 0x279d
	KindSP3S      = 0x947a

	RemoteIR       RemoteType = 0x26
	RemoteRF433Mhz RemoteType = 0xb2
	RemoteRF315Mhz RemoteType = 0xd7
)

type Device interface {
	Kind() int
	MAC() net.HardwareAddr
	Addr() *net.UDPAddr
	Interface() *net.UDPAddr
}

func NewDevice(kind int, mac net.HardwareAddr, addr, iface net.UDPAddr) Device {
	switch kind {
	case KindSP3S:
		return NewSP3S(mac, addr, iface)

	case KindRMProPlus:
		return NewRMProPlus(mac, addr, iface)

	case KindRM3Mini:
		return NewRM3Mini(mac, addr, iface)
	}

	return internal.NewDevice(kind, mac, addr, iface)
}