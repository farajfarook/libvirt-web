package main

import (
	"errors"

	libvirt "github.com/libvirt/libvirt-go"
)

// GetDomain by name
func (conn *Connection) GetDomain(name string) (*DomainInfo, error) {
	filters := libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE
	doms, err := conn.virt.ListAllDomains(filters)
	if err != nil {
		return nil, err
	}
	for _, dom := range doms {
		domName, _ := dom.GetName()
		if domName == name {
			dInfo := DomainInfo{}
			dInfo.Name = domName
			dInfo.StatusCode, _, _ = dom.GetState()
			dInfo.Status = decodeDomainState(dInfo.StatusCode)
			return &dInfo, nil
		}
	}
	return nil, errors.New("not found")
}

// DomainInfo Domain info model
type DomainInfo struct {
	Name       string              `json:"name"`
	StatusCode libvirt.DomainState `json:"statusCode"`
	Status     string              `json:"status"`
}
