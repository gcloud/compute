// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import ()

// The sizes offered by the compute service.
type Sizes struct{}

type Size struct {
	Id        string
	Name      string
	Ram       string
	Disk      string
	Bandwidth string
	Price     string
}

// List available sizes.
func (s *Sizes) List() {}

// Show size information for a given id.
func (s *Sizes) Show(id string) {}
