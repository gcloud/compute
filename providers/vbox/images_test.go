// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package vbox

import (
	"encoding/json"
	"testing"

	p "github.com/gcloud/compute/providers"
)

var ImageName = "GCloudImage"

func init() {
	servers := &Servers{}
	servers.Create(&p.Server{Name: ImageName})
}

func Test_ImagesCreate(t *testing.T) {
	images := &Images{}
	result, err := images.Create(&p.Image{Name: ImageName})
	if err != nil {
		t.Error("Images Create failed with " + err.Error() + ".")
	}
	if result == nil {
		t.Error("Results should not be nil.")
	}
	var image p.Image
	err = json.Unmarshal(result, &image)
	if err != nil {
		t.Error("Images Create failed with " + err.Error() + ".")
	}
	if len(image.Id) < 40 {
		t.Error("Wrong value for id.")
	}
	if ImageName != image.Name {
		t.Errorf("Expected %s, but is %s", ImageName, image.Name)
	}
}

func Test_ImagesList(t *testing.T) {
	images := &Images{}
	results, err := images.List()
	if err != nil {
		t.Error("Images List failed with " + err.Error() + "(bool, error).")
	}
	if results == nil {
		t.Error("Results should not be nil.")
	}
	var response []p.Image
	err = json.Unmarshal(results, &response)

	if err != nil {
		t.Error("Images List failed with " + err.Error() + "(bool, error).")
	}
	for _, image := range response {
		if len(image.Id) < 40 {
			t.Error("Wrong value for id.")
		}
		if ImageName != image.Name {
			t.Errorf("Expected %s, but is %s", ImageName, image.Name)
		}
	}
}

func Test_ImagesShow(t *testing.T) {
	images := &Images{}
	result, err := images.Show(ImageName)
	if err != nil {
		t.Error("Images Show failed with " + err.Error() + ".")
	}
	if result == nil {
		t.Error("Results should not be nil.")
	}
	var image p.Image
	err = json.Unmarshal(result, &image)
	if err != nil {
		t.Error("Images Show failed with " + err.Error() + ".")
	}
	if len(image.Id) < 40 {
		t.Error("Wrong value for id.")
	}
	if ImageName != image.Name {
		t.Errorf("Expected %s, but is %s", ImageName, image.Name)
	}
}

func Test_ImagesDestroy(t *testing.T) {
	images := &Images{}
	ok, err := images.Destroy(ImageName)
	if !ok {
		t.Error("Images Destroy failed.")
	}
	if err != nil {
		t.Error("Images Destroy failed with " + err.Error() + ".")
	}
}

func Test_Done(t *testing.T) {
	servers := &Servers{}
	servers.Destroy(ImageName)
}
