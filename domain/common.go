package domain

import (
	"errors"

	libvirt "github.com/libvirt/libvirt-go"
)

func decodeDomainState(status libvirt.DomainState) string {
	switch status {
	case libvirt.DOMAIN_NOSTATE:
		return "nostate"
	case libvirt.DOMAIN_BLOCKED:
		return "blocked"
	case libvirt.DOMAIN_RUNNING:
		return "running"
	case libvirt.DOMAIN_PAUSED:
		return "paused"
	case libvirt.DOMAIN_SHUTDOWN:
		return "shutdown"
	case libvirt.DOMAIN_CRASHED:
		return "crashed"
	case libvirt.DOMAIN_PMSUSPENDED:
		return "pmsuspended"
	case libvirt.DOMAIN_SHUTOFF:
		return "shutoff"
	}
	return "unknown"
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

func findAll() ([]libvirt.Domain, error) {
	filters := libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE
	return conn.ListAllDomains(filters)
}
