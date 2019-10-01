package domain

import libvirt "github.com/libvirt/libvirt-go"

// List List all domains
func List() ([]SummaryInfo, error) {
	filters := libvirt.CONNECT_LIST_NETWORKS_INACTIVE | libvirt.CONNECT_LIST_NETWORKS_ACTIVE
	nets, err := conn.ListAllNetworks(filters)
	if err != nil {
		return nil, err
	}
	var infos []SummaryInfo
	for _, net := range nets {
		info := SummaryInfo{}
		info.Name, _ = net.GetName()
		info.Active, _ = net.IsActive()
		infos = append(infos, info)
	}
	return infos, nil
}

// SummaryInfo summary info model
type SummaryInfo struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}
