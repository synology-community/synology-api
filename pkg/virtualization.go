package client

import "github.com/synology-community/synology-api/pkg/api/virtualization"

type virtualizationClient struct {
	client *APIClient
}

// GetGuest implements virtualization.VirtualizationAPI.
func (v *virtualizationClient) GetGuest(name string) (*virtualization.Guest, error) {
	return Get[virtualization.GetGuest, virtualization.Guest](v.client, &virtualization.GetGuest{Name: name}, virtualization.API_METHODS["GetGuest"])
}

// ListGuests implements virtualization.VirtualizationAPI.
func (v *virtualizationClient) ListGuests() (*virtualization.GuestList, error) {
	return List[virtualization.GuestList](v.client, virtualization.API_METHODS["ListGuests"])
}

func NewVirtualizationClient(client *APIClient) virtualization.VirtualizationAPI {
	return &virtualizationClient{client: client}
}