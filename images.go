// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	"github.com/gcloud/identity"
)

// The images available from the compute service.
type Images struct {
	identity.Account
}

type Image struct {
	Id   string
	Name string
}

// List images available to the account.
func (s *Images) List() {}

// Show image information for a given id.
func (s *Images) Show(id string) {}

// Create an image.
func (s *Images) Create() {}

// Destroy an image.
func (s *Images) Destroy() {}

// Distribute images to mutliple regions.
func (s *Images) Distribute() {}
