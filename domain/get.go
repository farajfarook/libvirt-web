package domain

import (
	libvirt "github.com/libvirt/libvirt-go"
)

func get(name string) (Info, error) {
	dom, err := find(name)
	if err != nil {
		return Info{}, err
	}
	dInfo := Info{}
	dInfo.Name = name
	dInfo.StatusCode, _, _ = dom.GetState()
	dInfo.Status = decodeDomainState(dInfo.StatusCode)
	return dInfo, nil
}

func getXML(name string) (string, error) {
	dom, err := find(name)
	if err != nil {
		return "", err
	}
	return dom.GetXMLDesc(libvirt.DOMAIN_XML_MIGRATABLE)
}

// Info domain info model
type Info struct {
	Name       string              `json:"name"`
	StatusCode libvirt.DomainState `json:"statusCode"`
	Status     string              `json:"status"`
}
