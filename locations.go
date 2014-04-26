// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	i "github.com/gcloud/identity"
)

type Locations interface {
	List() ([]Location, error)
	Show(string) (Location, error)
}

type Location interface {
	Id() string
	Name() string
	Region() string
}

func RegisterLocations(name string, locations Locations) {
	GetProvider(name).Locations = locations
}

func GetLocations(provider string, account *i.Account) Locations {
	Provider := GetProvider(provider)
	Provider.SetAccount(account)
	return Provider.Locations
}
