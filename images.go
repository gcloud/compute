// GCloud - Go Packages for Cloud Servicei.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	i "github.com/gcloud/identity"
)

type Images interface {
	New(Map) Image
	List() ([]Image, error)
	Show(Image) (Image, error)
	Create(Image) (Image, error)
	Destroy(Image) (bool, error)
}

type Image interface {
	Id() string
	Name() string
	Path() string
	String() string
	MarshalJSON() ([]byte, error)
	Map() Map
}

func RegisterImages(name string, images Images) {
	GetProvider(name).Images = images
}

func GetImages(provider string, account *i.Account) Images {
	Provider := GetProvider(provider)
	Provider.SetAccount(account)
	return Provider.Images
}
