package network

func list() ([]SummaryInfo, error) {
	nets, err := findAll()
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
