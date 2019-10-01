package domain

import libvirt "github.com/libvirt/libvirt-go"

// List List all domains
func List() ([]SummaryInfo, error) {
	doms, err := findAll()
	if err != nil {
		return nil, err
	}
	var dInfos []SummaryInfo
	for _, dom := range doms {
		dInfo := SummaryInfo{}
		dInfo.Name, _ = dom.GetName()
		dInfo.StatusCode, _, _ = dom.GetState()
		dInfo.Status = decodeDomainState(dInfo.StatusCode)
		dInfos = append(dInfos, dInfo)
	}
	return dInfos, nil
}

// SummaryInfo domain summary info model
type SummaryInfo struct {
	Name       string              `json:"name"`
	StatusCode libvirt.DomainState `json:"statusCode"`
	Status     string              `json:"status"`
}
