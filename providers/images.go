// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package providers

type Images interface {
	List() ([]Image, error)
	Show(string) (Image, error)
	Create(interface{}) (Image, error)
	Destroy(string) (bool, error)
}

type Image interface {
	Id() string
	Name() string
	File() string
}

func RegisterImages(name string, images Images) {
	GetProvider(name).Images = images
}
