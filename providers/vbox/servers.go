// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package vbox

import (
	p "github.com/gcloud/compute/providers"
)

type Servers struct{}

// List servers available on the account.
func (s *Servers) List() ([]byte, error) {
	return []byte(`[{}]`), nil
}

// Show server information for a given id.
func (s *Servers) Show(id string) ([]byte, error) {
	return []byte(`{}`), nil
}

// Create a server.
func (s *Servers) Create(n interface{}) ([]byte, error) {
	return []byte(`{}`), nil
}

// Destroy a server.
func (s *Servers) Destroy(id string) (bool, error) {
	return false, nil
}

// Reboot a server.
func (s *Servers) Reboot(id string) (bool, error) {
	return false, nil
}

// Start a server that is stopped.
func (s *Servers) Start(id string) (bool, error) {
	return false, nil
}

// Stop a server that is running.
func (s *Servers) Stop(id string) (bool, error) {
	return false, nil
}

func init() {
	p.RegisterServers("vbox", &Servers{})
}
