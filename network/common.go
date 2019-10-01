package network

import (
	"errors"

	libvirt "github.com/libvirt/libvirt-go"
)

func find(name string) (libvirt.Network, error) {
	nets, err := findAll()
	if err != nil {
		return libvirt.Network{}, err
	}
	for _, net := range nets {
		currName, _ := net.GetName()
		if currName == name {
			return net, nil
		}
	}
	return libvirt.Network{}, errors.New("not found")
}

func findAll() ([]libvirt.Network, error) {
	filters := libvirt.CONNECT_LIST_NETWORKS_INACTIVE | libvirt.CONNECT_LIST_NETWORKS_ACTIVE
	return conn.ListAllNetworks(filters)
}
