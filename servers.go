// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	"github.com/gcloud/identity"
	p "github.com/gcloud/providers"
)

// The Servers type interacts with Compute services.
type Servers struct {
	Account  *identity.Account
	Provider string
}

func (s *Servers) NewServer(m p.Map) p.Server {
	provider := p.GetProvider(s.Provider)
	return provider.Servers.NewServer(m)
}

// List servers available on the account.
func (s *Servers) List() ([]p.Server, error) {
	provider := p.GetProvider(s.Provider)
	provider.SetAccount(s.Account)
	return provider.Servers.List()
}

// Show server information for a given id.
func (s *Servers) Show(id string) (p.Server, error) {
	provider := p.GetProvider(s.Provider)
	provider.SetAccount(s.Account)
	return provider.Servers.Show(id)
}

// Create a server.
func (s *Servers) Create(n interface{}) (p.Server, error) {
	provider := p.GetProvider(s.Provider)
	provider.SetAccount(s.Account)
	return provider.Servers.Create(n)
}

// Destroy a server.
func (s *Servers) Destroy(id string) (bool, error) {
	provider := p.GetProvider(s.Provider)
	provider.SetAccount(s.Account)
	ok, err := provider.Servers.Destroy(id)
	return ok, err
}

// Reboot a server.
func (s *Servers) Reboot(id string) (bool, error) {
	provider := p.GetProvider(s.Provider)
	provider.SetAccount(s.Account)
	ok, err := provider.Servers.Reboot(id)
	return ok, err
}

// Start a server that is stopped.
func (s *Servers) Start(id string) (bool, error) {
	provider := p.GetProvider(s.Provider)
	provider.SetAccount(s.Account)
	ok, err := provider.Servers.Start(id)
	return ok, err
}

// Stop a server that is running.
func (s *Servers) Stop(id string) (bool, error) {
	provider := p.GetProvider(s.Provider)
	provider.SetAccount(s.Account)
	ok, err := provider.Servers.Stop(id)
	return ok, err
}
