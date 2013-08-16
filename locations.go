// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import ()

// The locations provided by a compute service.
type Locations struct{}

type Location struct {
	Id      string
	Name    string
	Country string
}

// List available locations.
func (s *Locations) List() {}

// Show location information for a given id.
func (s *Locations) Show(id string) {}
