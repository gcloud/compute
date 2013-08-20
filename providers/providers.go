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
	List() string
	Show(string)
	Create()
	Destroy()
}

type Locations interface {
	List() string
	Show(string)
}

type Sizes interface {
	List() string
	Show(string)
}

var Providers = make(map[string]*Provider)

func RegisterServers(name string, servers Servers) {
	if servers == nil {
		panic("compute: Servers is nil.")
	}
	if _, ok := Providers[name]; !ok {
		Providers[name] = &Provider{}
	}
	Providers[name].Servers = servers
}
