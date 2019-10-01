package domain

// create godoc
// @Tags domain
// @Param request body domain.CreateRequest true "Create Request"
// @Success 200 {object} domain.CreateResponse
// @Router /domains [post]
func create(request CreateRequest) (CreateResponse, error) {
	return CreateResponse{Name: request.Name}, nil
}

//CreateRequest dogoc
type CreateRequest struct {
	Name    string `json:"name"`
	Memory  int    `json:"memory"`
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
