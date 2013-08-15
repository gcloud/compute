// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package vbox

import (
	p "github.com/gcloud/compute/providers"
)

type Servers struct{}

// List servers available on the account.
func (s *Servers) List() string {
	return "virtualbox"
}

// Show server information for a given id.
func (s *Servers) Show(id string) {}

// Create a server.
func (s *Servers) Create() {}

// Destroy a server.
func (s *Servers) Destroy() {}

// Reboot a server.
func (s *Servers) Reboot() {}

// Start a server that is stopped.
func (s *Servers) Start() {}

// Stop a server that is running.
func (s *Servers) Stop() {}

func init() {
	p.RegisterServers("vbox", &Servers{})
}
