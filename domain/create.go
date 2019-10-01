package domain

import (
	"github.com/labstack/echo"
	libvirt "github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

// create godoc
// @Tags domain
// @Param request body domain.CreateRequest true "Create Request"
// @Success 200 {object} domain.CreateResponse
// @Router /domains [post]
func create(request CreateRequest) (CreateResponse, error) {
	domcfg := &libvirtxml.Domain{
		Type: request.Type,
		Name: request.Name,
		Memory: &libvirtxml.DomainMemory{
			Value:    request.Memory,
			Unit:     "MB",
			DumpCore: "on",
		},
		VCPU: &libvirtxml.DomainVCPU{
			Value: request.VCPU,
		},
		CPU: &libvirtxml.DomainCPU{
			Mode: "host-model",
		},
		OS: &libvirtxml.DomainOS{
			Type: &libvirtxml.DomainOSType{
				Arch:    request.Arch,
				Machine: "pc",
				Type:    "hvm",
			},
		},
		Devices: &libvirtxml.DomainDeviceList{
			Disks: []libvirtxml.DomainDisk{},
		},
	}
	xml, _ := domcfg.Marshal()
	dom, err := conn.DomainCreateXML(xml, libvirt.DOMAIN_NONE)
	if err != nil {
		return CreateResponse{}, echo.NewHTTPError(500, err.Error())
	}
	resp := CreateResponse{}
	resp.Name, _ = dom.GetName()
	return resp, nil
}

//CreateRequest dogoc
type CreateRequest struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Memory  uint   `json:"memory"`
	VCPU    int    `json:"vcpu"`
	Arch    string `json:"arch"`
	CDROM   string `json:"cdrom"`
	Disk    string `json:"disk"`
	Network string `json:"network"`
}

//CreateResponse godoc
type CreateResponse struct {
	Name string `json:"name"`
}
