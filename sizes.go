// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	i "github.com/gcloud/identity"
)

type Sizes interface {
	List() ([]Size, error)
	Show(string) (Size, error)
}

type Size interface {
	Id() string
	Name() string
}

func RegisterSizes(name string, sizes Sizes) {
	GetProvider(name).Sizes = sizes
}

func GetSizes(provider string, account *i.Account) Sizes {
	Provider := GetProvider(provider)
	Provider.SetAccount(account)
	return Provider.Sizes
}
