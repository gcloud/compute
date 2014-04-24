// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	p "github.com/gcloud/compute/providers"
	"github.com/gcloud/identity"
)

// The locations provided by a compute service.
type Locations struct {
	Account  *identity.Account
	Provider string
}

// List available locations.
func (l *Locations) List() ([]p.Location, error) {
	provider := p.GetProvider(l.Provider)
	provider.SetAccount(l.Account)
	return provider.Locations.List()
}

// Show location information for a given id.
func (l *Locations) Show(id string) (p.Location, error) {
	provider := p.GetProvider(l.Provider)
	provider.SetAccount(l.Account)
	return provider.Locations.Show(id)
}
