// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package providers

type Servers interface {
	NewServer(Map) Server
	List() ([]Server, error)
	Show(string) (Server, error)
	Create(interface{}) (Server, error)
	Destroy(string) (bool, error)
	Reboot(string) (bool, error)
	Start(string) (bool, error)
	Stop(string) (bool, error)
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
}

func RegisterServers(name string, servers Servers) {
	GetProvider(name).Servers = servers
}
