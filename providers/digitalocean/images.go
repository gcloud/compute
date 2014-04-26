// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package digitalocean

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gcloud/compute"
	"github.com/mitchellh/mapstructure"
)

func init() {
	name := "digitalocean"
	provider := compute.GetProvider(name)
	provider.Endpoint = "https://api.digitalocean.com"
	compute.RegisterImages(name, &Images{provider: provider})
}

type iResult struct {
	Status        string
	Image         *image
	Error_message string `json:",omitempty"`
}

type iResults struct {
	Status        string
	Images        []*image
	Error_message string `json:",omitempty"`
}

type image struct {
	Id           int
	Name         string
	Distribution string
	Slug         string
	Public       bool
}

func (i *image) toImage() *Image {
	return &Image{data: i}
}

type Image struct {
	data    *image
	generic compute.Map
}

func (i *Image) Id() string {
	return fmt.Sprintf("%d", i.data.Id)
}
func (i *Image) Name() string {
	return i.data.Name
}
func (i *Image) Path() string {
	return ""
}
func (i *Image) String() string {
	b, err := i.MarshalJSON()
	if err != nil {
		return ""
	}
	return string(b)
}
func (i *Image) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.data)
}
func (i *Image) Map() compute.Map {
	return i.generic
}

type Images struct {
	provider *compute.Provider
}

func (i *Images) New(m compute.Map) compute.Image {
	var newImage image
	err := mapstructure.Decode(m, &newImage)
	if err != nil {
		return nil
	}
	im := newImage.toImage()
	im.generic = m
	return im
}

// List images available on the account.
func (i *Images) List() ([]compute.Image, error) {
	var r []compute.Image
	response, err := request(i.provider, "GET", "/images", nil)
	if err != nil {
		return r, err
	}
	var results iResults
	err = json.Unmarshal(response, &results)
	if err != nil {
		return r, err
	}
	if results.Status != "OK" {
		return r, errors.New(results.Error_message)
	}
	images := make([]compute.Image, 0)
	for _, v := range results.Images {
		images = append(images, v.toImage())
	}
	return images, nil
}

// Show image information for a given id.
func (i *Images) Show(image compute.Image) (compute.Image, error) {
	var r compute.Image
	response, err := request(i.provider, "GET", "/images/"+image.Id(), nil)
	if err != nil {
		return r, err
	}
	var result iResult
	err = json.Unmarshal(response, &result)
	if err != nil {
		return r, err
	}
	if result.Status != "OK" {
		return r, errors.New(result.Error_message)
	}
	return result.Image.toImage(), nil
}

// Create a image.
func (i *Images) Create(image compute.Image) (compute.Image, error) {
	//GET https://api.digitalocean.com/droplets/[droplet_id]/snapshot/?name=[snapshot_name]
	return image, nil
}

// Delete a image.
func (i *Images) Destroy(image compute.Image) (bool, error) {
	return event(i.provider, "/images/%s/destroy", image.Id())
}
