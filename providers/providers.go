// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package providers

import (
	"github.com/gcloud/identity"
)

type Provider struct {
	Account   identity.Account
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
	Id   string
	Name string
}

type Locations interface {
	List() string
	Show(string)
}

type Sizes interface {
	List() string
	Show(string)
}

var providers = make(map[string]*Provider)

func GetProvider(name string, account identity.Account) *Provider {
	if _, ok := providers[name]; !ok {
		providers[name] = &Provider{}
	}
	providers[name].Account = account
	return providers[name]
}

func RegisterServers(name string, servers Servers) {
	if servers == nil {
		panic("compute: Images is nil.")
	}
	if _, ok := providers[name]; !ok {
		providers[name] = &Provider{}
	}
	providers[name].Servers = servers
}

func RegisterImages(name string, images Images) {
	if images == nil {
		panic("compute: Images is nil.")
	}
	if _, ok := providers[name]; !ok {
		providers[name] = &Provider{}
	}
	providers[name].Images = images
}

func RegisterLocations(name string, locations Locations) {
	if locations == nil {
		panic("compute: Images is nil.")
	}
	if _, ok := providers[name]; !ok {
		providers[name] = &Provider{}
	}
	providers[name].Locations = locations
}

func RegisterSizes(name string, sizes Sizes) {
	if sizes == nil {
		panic("compute: Images is nil.")
	}
	if _, ok := providers[name]; !ok {
		providers[name] = &Provider{}
	}
	providers[name].Sizes = sizes
}
