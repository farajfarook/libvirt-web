package domain

import (
	"encoding/xml"
	"errors"

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

	err := createPreprocess(&request)
	if err != nil {
		return CreateResponse{}, echo.NewHTTPError(400, err.Error())
	}

	domcfg := &libvirtxml.Domain{
		Type: request.Type,
		Name: request.Name,
		Memory: &libvirtxml.DomainMemory{
			Value: request.Memory,
			Unit:  "MiB",
		},
		VCPU: &libvirtxml.DomainVCPU{
			Placement: "static",
			Value:     request.VCPU,
		},
		CPU: &libvirtxml.DomainCPU{},
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

func createPreprocess(request *CreateRequest) error {
	if request.Arch == "" {
		request.Arch, _ = getHostArchitecture()
	}
	if request.Type == "" {
		request.Type = "kvm"
	}
	if request.Memory == 0 {
		return errors.New("memory is empty")
	}
	if request.VCPU == 0 {
		return errors.New("vCPU is empty")
	}
	return nil
}

func getHostArchitecture() (string, error) {
	type HostCapabilities struct {
		XMLName xml.Name `xml:"capabilities"`
		Host    struct {
			XMLName xml.Name `xml:"host"`
			CPU     struct {
				XMLName xml.Name `xml:"cpu"`
				Arch    string   `xml:"arch"`
			}
		}
	}

	info, err := conn.GetCapabilities()
	if err != nil {
		return "", err
	}

	capabilities := HostCapabilities{}
	xml.Unmarshal([]byte(info), &capabilities)

	return capabilities.Host.CPU.Arch, nil
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
