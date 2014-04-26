// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	i "github.com/gcloud/identity"
)

type Map map[string]interface{}

type Provider struct {
	Endpoint  string
	Account   *i.Account
	Servers   Servers
	Images    Images
	Locations Locations
	Sizes     Sizes
}

func (p *Provider) SetAccount(account *i.Account) {
	p.Account = account
}

func (p *Provider) GetAccount() *i.Account {
	return p.Account
}

var providers = make(map[string]*Provider)

func GetProvider(name string) *Provider {
	if _, ok := providers[name]; !ok {
		providers[name] = &Provider{}
	}
	return providers[name]
}
