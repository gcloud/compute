// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package providers

type Provider struct {
	Servers Servers
	// Images() Images
	// Locations() Locations
	// Sizes() Sizes
}

type Servers interface {
	List() string
	// Show(string)
	// Create()
	// Destroy()
	// Reboot()
	// Start()
	// Stop()
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
