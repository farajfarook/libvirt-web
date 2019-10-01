package domain

//Create create domain
func Create(request CreateRequest) (*CreateResponse, error) {
	return nil, nil
}

//CreateRequest model
type CreateRequest struct {
	Name    string `json:"name"`
	Memory  int    `json:"memory"`
	VCPU    int    `json:"vcpu"`
	Arch    string `json:"arch"`
	CDROM   string `json:"cdrom"`
	Disk    string `json:"disk"`
	Network string `json:"network"`
}

//CreateResponse model
type CreateResponse struct {
	SummaryInfo
}
