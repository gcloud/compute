// GCloud - Go Packages for Cloud Servicei.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	"github.com/gcloud/identity"
	p "github.com/gcloud/providers"
)

// The images available from the compute service.
type Images struct {
	Account  *identity.Account
	Provider string
}

// List images available on the account.
func (i *Images) List() ([]p.Image, error) {
	provider := p.GetProvider(i.Provider, i.Account)
	return provider.Images.List()
}

// Show image information for a given id.
func (i *Images) Show(id string) (p.Image, error) {
	provider := p.GetProvider(i.Provider, i.Account)
	return provider.Images.Show(id)
}

// Create a image.
func (i *Images) Create(n interface{}) (p.Image, error) {
	provider := p.GetProvider(i.Provider, i.Account)
	return provider.Images.Create(n)
}

// Destroy a image.
func (i *Images) Destroy(id string) (bool, error) {
	provider := p.GetProvider(i.Provider, i.Account)
	return provider.Images.Destroy(id)
}
