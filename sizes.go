// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	p "github.com/gcloud/compute/providers"
	"github.com/gcloud/identity"
)

// The sizes offered by the compute service.
type Sizes struct {
	Account  *identity.Account
	Provider string
}

// List available sizes.
func (s *Sizes) List() ([]p.Size, error) {
	provider := p.GetProvider(s.Provider)
	provider.SetAccount(s.Account)
	return provider.Sizes.List()
}

// Show size information for a given id.
func (s *Sizes) Show(id string) (p.Size, error) {
	provider := p.GetProvider(s.Provider)
	provider.SetAccount(s.Account)
	return provider.Sizes.Show(id)
}
