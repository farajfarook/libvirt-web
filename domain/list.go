package domain

import libvirt "github.com/libvirt/libvirt-go"

// list godoc
// @Tags domain
// @Success 200 {object} domain.SummaryInfo
// @Router /domains [get]
func list() ([]SummaryInfo, error) {
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

// SummaryInfo godoc
type SummaryInfo struct {
	Name       string              `json:"name"`
	StatusCode libvirt.DomainState `json:"statusCode"`
	Status     string              `json:"status"`
}
