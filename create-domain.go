package main

//CreateDomain create domain
func (conn *Connection) CreateDomain(request CreateDomainRequest) (*CreateDomainResponse, error) {
	return nil, nil
}

//CreateDomainRequest model
type CreateDomainRequest struct {
	Name    string `json:"name"`
	Memory  int    `json:"memory"`
	VCPU    int    `json:"vcpu"`
	Arch    string `json:"arch"`
	CDROM   string `json:"cdrom"`
	Disk    string `json:"Disk"`
	Network string `json:"network"`
}

//CreateDomainResponse model
type CreateDomainResponse struct {
	DomainSummaryInfo
}
