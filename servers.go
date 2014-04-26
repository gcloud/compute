// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	i "github.com/gcloud/identity"
)

type Servers interface {
	New(Map) Server
	List() ([]Server, error)
	Show(Server) (Server, error)
	Create(Server) (Server, error)
	Destroy(Server) (bool, error)
	Reboot(Server) (bool, error)
	Start(Server) (bool, error)
	Stop(Server) (bool, error)
}

type Server interface {
	Id() string
	Name() string
	State() string
	Ips(string) []string
	Size() string
	Image() string
	String() string
	MarshalJSON() ([]byte, error)
	Map() Map
}

func RegisterServers(provider string, servers Servers) {
	GetProvider(provider).Servers = servers
}

func GetServers(provider string, account *i.Account) Servers {
	Provider := GetProvider(provider)
	Provider.SetAccount(account)
	return Provider.Servers
}
