// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package providers

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
