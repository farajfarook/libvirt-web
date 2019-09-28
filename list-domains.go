package main

import libvirt "github.com/libvirt/libvirt-go"

func listDomains(conn *libvirt.Connect) ([]domainSummaryInfo, error) {
	filters := libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE
	doms, err := conn.ListAllDomains(filters)
	if err != nil {
		return nil, err
	}
	dInfos := []domainSummaryInfo{}
	for _, dom := range doms {
		name, _ := dom.GetName()
		status, _, _ := dom.GetState()
		dInfos = append(dInfos, domainSummaryInfo{
			name:       name,
			statusCode: status,
			status:     decodeDomainState(status),
		})
	}
	return dInfos, nil
}

type domainSummaryInfo struct {
	name       string
	statusCode libvirt.DomainState
	status     string
}
