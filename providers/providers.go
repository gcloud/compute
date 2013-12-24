// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package providers

import (
	"github.com/gcloud/identity"
)

type Provider struct {
	Account   *identity.Account
	Servers   Servers
	Images    Images
	Locations Locations
	Sizes     Sizes
}

type Servers interface {
	List() ([]byte, error)
	Show(string) ([]byte, error)
	Create(*Server) ([]byte, error)
	Destroy(id string) (bool, error)
	Reboot(id string) (bool, error)
	Start(id string) (bool, error)
	Stop(id string) (bool, error)
}

type Server struct {
	Id         string
	Name       string
	State      string
	PublicIps  []string
	PrivateIps []string
	Size       string
	Image      string
}

type Images interface {
	List() ([]byte, error)
	Show(string) ([]byte, error)
	Create(*Image) ([]byte, error)
	Destroy(id string) (bool, error)
}

type Image struct {
	Id     string
	Name   string
	Source string
}

type Locations interface {
	List() ([]byte, error)
	Show(string) ([]byte, error)
}

type Location struct {
	Id     string
	Name   string
	Region string
}

type Sizes interface {
	List() ([]byte, error)
	Show(string) ([]byte, error)
}

type Size struct {
	Id        string
	Name      string
	Ram       string
	Disk      string
	Bandwidth string
	Price     string
}

var providers = make(map[string]*Provider)

func GetProvider(name string, account *identity.Account) *Provider {
	if _, ok := providers[name]; !ok {
		providers[name] = &Provider{}
	}
	if account != nil {
		providers[name].Account = account
	}
	return providers[name]
}

func RegisterServers(name string, servers Servers) {
	GetProvider(name, nil).Servers = servers
}

func RegisterImages(name string, images Images) {
	GetProvider(name, nil).Images = images
}

func RegisterLocations(name string, locations Locations) {
	GetProvider(name, nil).Locations = locations
}

func RegisterSizes(name string, sizes Sizes) {
	GetProvider(name, nil).Sizes = sizes
}
