// GCloud - Go Packages for Cloud Servicei.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	"encoding/json"

	p "github.com/gcloud/compute/providers"
	"github.com/gcloud/identity"
)

// The images available from the compute service.
type Images struct {
	Account  *identity.Account
	Provider string
}

// List images available on the account.
func (i *Images) List() (*[]p.Image, error) {
	provider := p.GetProvider(i.Provider, i.Account)
	result, err := provider.Images.List()
	if err != nil {
		return nil, err
	}
	var records []p.Image
	err = json.Unmarshal(result, &records)

	if err != nil {
		return nil, err
	}
	return &records, err
}

// Show image information for a given id.
func (i *Images) Show(id string) (*p.Image, error) {
	provider := p.GetProvider(i.Provider, i.Account)
	result, err := provider.Images.Show(id)
	if err != nil {
		return nil, err
	}
	var record p.Image
	err = json.Unmarshal(result, &record)

	if err != nil {
		return nil, err
	}
	return &record, err
}

// Create a image.
func (i *Images) Create(n *p.Image) (*p.Image, error) {
	provider := p.GetProvider(i.Provider, i.Account)
	result, err := provider.Images.Create(n)
	if err != nil {
		return nil, err
	}
	var record p.Image
	err = json.Unmarshal(result, &record)

	if err != nil {
		return nil, err
	}
	return &record, err
}

// Destroy a image.
func (i *Images) Destroy(id string) (bool, error) {
	provider := p.GetProvider(i.Provider, i.Account)
	ok, err := provider.Images.Destroy(id)
	return ok, err
}
