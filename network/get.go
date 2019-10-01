package network

import (
	"github.com/labstack/echo"
	libvirt "github.com/libvirt/libvirt-go"
)

// get godoc
// @Tags network
// @Param name path string true "Name"
// @Success 200 {object} network.Info
// @Router /networks/{name} [get]
func get(name string) (Info, error) {
	net, err := find(name)
	if err != nil {
		return Info{}, echo.NewHTTPError(404, err.Error())
	}
	info := Info{}
	info.Name = name
	info.Active, _ = net.IsActive()
	info.Persistent, _ = net.IsPersistent()
	info.Bridge, _ = net.GetBridgeName()
	return info, nil
}

// getXML godoc
// @Tags network
// @Param name path string true "Name"
// @Success 200 {object} string
// @Router /networks/{name}/_xml [get]
func getXML(name string) (string, error) {
	dom, err := find(name)
	if err != nil {
		return "", echo.NewHTTPError(404, err.Error())
	}
	return dom.GetXMLDesc(libvirt.NETWORK_XML_INACTIVE)
}

// Info godoc
type Info struct {
	Name       string `json:"name"`
	Active     bool   `json:"active"`
	Persistent bool   `json:"persistent"`
	Bridge     string `json:"bridge"`
}
