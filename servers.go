// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	"encoding/json"
	p "github.com/gcloud/compute/providers"
	"github.com/gcloud/identity"
)

// The Servers type interacts with Compute services.
type Servers struct {
	Account  identity.Account
	Provider string
}

// List servers available on the account.
func (s *Servers) List() (*[]p.Server, error) {
	provider := p.GetProvider(s.Provider, s.Account)
	result, err := provider.Servers.List()
	if err != nil {
		return nil, err
	}
	var records []p.Server
	err = json.Unmarshal(result, &records)

	if err != nil {
		return nil, err
	}
	return &records, err
}

// Show server information for a given id.
func (s *Servers) Show(id string) (*p.Server, error) {
	provider := p.GetProvider(s.Provider, s.Account)
	result, err := provider.Servers.Show(id)
	if err != nil {
		return nil, err
	}
	var record p.Server
	err = json.Unmarshal(result, &record)

	if err != nil {
		return nil, err
	}
	return &record, err
}

// Create a server.
func (s *Servers) Create(n *p.Server) (*p.Server, error) {
	provider := p.GetProvider(s.Provider, s.Account)
	result, err := provider.Servers.Create(n)
	if err != nil {
		return nil, err
	}
	var record p.Server
	err = json.Unmarshal(result, &record)

	if err != nil {
		return nil, err
	}
	return &record, err
}

// Destroy a server.
func (s *Servers) Destroy(id string) (bool, error) {
	provider := p.GetProvider(s.Provider, s.Account)
	ok, err := provider.Servers.Destroy(id)
	return ok, err
}

// Reboot a server.
func (s *Servers) Reboot(id string) (bool, error) {
	provider := p.GetProvider(s.Provider, s.Account)
	ok, err := provider.Servers.Reboot(id)
	return ok, err
}

// Start a server that is stopped.
func (s *Servers) Start(id string) (bool, error) {
	provider := p.GetProvider(s.Provider, s.Account)
	ok, err := provider.Servers.Start(id)
	return ok, err
}

// Stop a server that is running.
func (s *Servers) Stop(id string) (bool, error) {
	provider := p.GetProvider(s.Provider, s.Account)
	ok, err := provider.Servers.Stop(id)
	return ok, err
}
