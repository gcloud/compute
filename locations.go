// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	"encoding/json"

	"github.com/gcloud/identity"
	p "github.com/gcloud/providers"
)

// The locations provided by a compute service.
type Locations struct {
	Account  *identity.Account
	Provider string
}

// List available locations.
func (l *Locations) List() (*[]p.Location, error) {
	provider := p.GetProvider(l.Provider, l.Account)
	result, err := provider.Locations.List()
	if err != nil {
		return nil, err
	}
	var records []p.Location
	err = json.Unmarshal(result, &records)

	if err != nil {
		return nil, err
	}
	return &records, err
}

// Show location information for a given id.
func (l *Locations) Show(id string) (*p.Location, error) {
	provider := p.GetProvider(l.Provider, l.Account)
	result, err := provider.Locations.Show(id)
	if err != nil {
		return nil, err
	}
	var record p.Location
	err = json.Unmarshal(result, &record)

	if err != nil {
		return nil, err
	}
	return &record, err
}
