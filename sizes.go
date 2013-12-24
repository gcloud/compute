// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	"encoding/json"

	p "github.com/gcloud/compute/providers"
	"github.com/gcloud/identity"
)

// The sizes offered by the compute service.
type Sizes struct {
	Account  *identity.Account
	Provider string
}

// List available sizes.
func (s *Sizes) List() (*[]p.Size, error) {
	provider := p.GetProvider(s.Provider, s.Account)
	result, err := provider.Sizes.List()
	if err != nil {
		return nil, err
	}
	var records []p.Size
	err = json.Unmarshal(result, &records)

	if err != nil {
		return nil, err
	}
	return &records, err
}

// Show size information for a given id.
func (s *Sizes) Show(id string) (*p.Size, error) {
	provider := p.GetProvider(s.Provider, s.Account)
	result, err := provider.Sizes.Show(id)
	if err != nil {
		return nil, err
	}
	var record p.Size
	err = json.Unmarshal(result, &record)

	if err != nil {
		return nil, err
	}
	return &record, err
}
