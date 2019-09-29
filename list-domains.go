package main

import libvirt "github.com/libvirt/libvirt-go"

// ListDomains List all domains
func (conn *Connection) ListDomains() ([]DomainSummaryInfo, error) {
	filters := libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE
	doms, err := conn.virt.ListAllDomains(filters)
	if err != nil {
		return nil, err
	}
	var dInfos []DomainSummaryInfo
	for _, dom := range doms {
		dInfo := DomainSummaryInfo{}
		dInfo.Name, _ = dom.GetName()
		dInfo.StatusCode, _, _ = dom.GetState()
		dInfo.Status = decodeDomainState(dInfo.StatusCode)
		dInfos = append(dInfos, dInfo)
	}
	return dInfos, nil
}

// DomainSummaryInfo domain summary info model
type DomainSummaryInfo struct {
	Name       string              `json:"name"`
	StatusCode libvirt.DomainState `json:"statusCode"`
	Status     string              `json:"status"`
}
