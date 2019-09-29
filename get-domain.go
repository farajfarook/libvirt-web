package main

import (
	"errors"

	libvirt "github.com/libvirt/libvirt-go"
)

// GetDomain by name
func (conn *Connection) GetDomain(name string) (DomainInfo, error) {
	dom, err := conn.getDomain(name)
	if err != nil {
		return DomainInfo{}, err
	}
	dInfo := DomainInfo{}
	dInfo.Name = name
	dInfo.StatusCode, _, _ = dom.GetState()
	dInfo.Status = decodeDomainState(dInfo.StatusCode)
	return dInfo, nil
}

//GetDomainXML get XML content of domain
func (conn *Connection) GetDomainXML(name string) (string, error) {
	dom, err := conn.getDomain(name)
	if err != nil {
		return "", err
	}
	return dom.GetXMLDesc(libvirt.DOMAIN_XML_MIGRATABLE)
}

func (conn *Connection) getDomain(name string) (libvirt.Domain, error) {
	filters := libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE
	doms, err := conn.virt.ListAllDomains(filters)
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

// DomainInfo domain info model
type DomainInfo struct {
	Name       string              `json:"name"`
	StatusCode libvirt.DomainState `json:"statusCode"`
	Status     string              `json:"status"`
}
