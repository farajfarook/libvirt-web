package domain

import (
	"errors"

	libvirt "github.com/libvirt/libvirt-go"
)

// Get by name
func Get(name string) (Info, error) {
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

//GetXML get XML content of domain
func GetXML(name string) (string, error) {
	dom, err := find(name)
	if err != nil {
		return "", err
	}
	return dom.GetXMLDesc(libvirt.DOMAIN_XML_MIGRATABLE)
}

func find(name string) (libvirt.Domain, error) {
	filters := libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE
	doms, err := conn.ListAllDomains(filters)
	if err != nil {
		return libvirt.Domain{}, err
	}
	for _, dom := range doms {
		domName, _ := dom.GetName()
		if domName == name {
			return dom, nil
		}
	}
	return libvirt.Domain{}, errors.New("not found")
}

// Info domain info model
type Info struct {
	Name       string              `json:"name"`
	StatusCode libvirt.DomainState `json:"statusCode"`
	Status     string              `json:"status"`
}
