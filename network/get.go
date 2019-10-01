package network

import (
	libvirt "github.com/libvirt/libvirt-go"
)

func get(name string) (Info, error) {
	net, err := find(name)
	if err != nil {
		return Info{}, err
	}
	info := Info{}
	info.Name = name
	info.Active, _ = net.IsActive()
	info.Persistent, _ = net.IsPersistent()
	return info, nil
}

func getXML(name string) (string, error) {
	dom, err := find(name)
	if err != nil {
		return "", err
	}
	return dom.GetXMLDesc(libvirt.NETWORK_XML_INACTIVE)
}

// Info info model
type Info struct {
	Name       string `json:"name"`
	Active     bool   `json:"active"`
	Persistent bool   `json:"persistent"`
}
