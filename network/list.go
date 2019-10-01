package network

import "github.com/labstack/echo"

// list godoc
// @Tags network
// @Success 200 {object} network.SummaryInfo
// @Router /networks [get]
func list() ([]SummaryInfo, error) {
	nets, err := findAll()
	if err != nil {
		return nil, echo.NewHTTPError(500, err.Error())
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

// SummaryInfo godoc
type SummaryInfo struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}
