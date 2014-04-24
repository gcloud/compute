package providers

import (
	"github.com/gcloud/identity"
)

type Map map[string]interface{}

type Provider struct {
	Endpoint  string
	Account   *identity.Account
	Servers   Servers
	Images    Images
	Locations Locations
	Sizes     Sizes
}

func (p *Provider) SetAccount(account *identity.Account) {
	p.Account = account
}

func (p *Provider) GetAccount() *identity.Account {
	return p.Account
}

var providers = make(map[string]*Provider)

func GetProvider(name string) *Provider {
	if _, ok := providers[name]; !ok {
		providers[name] = &Provider{}
	}
	return providers[name]
}
