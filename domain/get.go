package domain

import (
	"github.com/labstack/echo"
	libvirt "github.com/libvirt/libvirt-go"
)

// get godoc
// @Tags domain
// @Param name path string true "Name"
// @Success 200 {object} domain.Info
// @Router /domains/{name} [get]
func get(name string) (Info, error) {
	dom, err := find(name)
	if err != nil {
		return Info{}, echo.NewHTTPError(404, err.Error())
	}
	dInfo := Info{}
	dInfo.Name = name
	dInfo.StatusCode, _, _ = dom.GetState()
	dInfo.Status = decodeDomainState(dInfo.StatusCode)
	return dInfo, nil
}

// getXML godoc
// @Tags domain
// @Param name path string true "Name"
// @Success 200 {object} string
// @Router /domains/{name}/_xml [get]
func getXML(name string) (string, error) {
	dom, err := find(name)
	if err != nil {
		return "", echo.NewHTTPError(404, err.Error())
	}
	return dom.GetXMLDesc(libvirt.DOMAIN_XML_MIGRATABLE)
}

// Info godoc
type Info struct {
	Name       string              `json:"name"`
	StatusCode libvirt.DomainState `json:"statusCode"`
	Status     string              `json:"status"`
}
