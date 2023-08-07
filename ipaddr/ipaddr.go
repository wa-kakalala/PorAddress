package ipaddr

import (
	"errors"
	"net"
)

type AddressList struct {
	NetNames  []string
	Addresses []string
}

func (al *AddressList) GetAddressList() error {
	interfaces, err := net.Interfaces()
	if err != nil {
		return err
	}
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
				al.NetNames = append(al.NetNames, iface.Name)
				al.Addresses = append(al.Addresses, ipNet.IP.String())
			}
		}
	}

	if len(al.Addresses) == 0 {
		return errors.New("get address err")
	} else {
		return nil
	}

}
