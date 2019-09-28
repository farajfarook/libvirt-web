package main

import libvirt "github.com/libvirt/libvirt-go"

func listDomains() ([]domainSummaryInfo, error) {
	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE)
	if err != nil {
		return nil, err
	}
	dInfos := []domainSummaryInfo{}
	for _, dom := range doms {
		name, _ := dom.GetName()
		status, _, _ := dom.GetState()
		dInfos = append(dInfos, domainSummaryInfo{
			name:   name,
			status: status,
		})
	}
	return dInfos, nil
}

type domainSummaryInfo struct {
	name   string
	status libvirt.DomainState
}
