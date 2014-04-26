// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package vbox

import (
	"testing"

	"github.com/gcloud/compute"
)

var TestImage = &Image{name: "GCloudImage"}
var TestImageServer compute.Server

func init() {
	servers := &Servers{}
	s, err := servers.Create(&Server{name: TestImage.Name()})
	if err != nil {
		println("Server not created.")
		return
	}
	TestImageServer = s
}

func Test_ImagesCreate(t *testing.T) {
	images := &Images{}
	result, err := images.Create(TestImage)
	if err != nil {
		t.Errorf("Images Create failed with %s.", err)
	}
	if result == nil {
		t.Error("Images Create result should not be nil.")
		return
	}
	if len(result.Id()) < 40 {
		t.Error("Wrong value for id.")
	}
	if TestImage.Name() != result.Name() {
		t.Errorf("Expected %s, but is %s", TestImage.Name(), result.Name())
	}
}

func Test_ImagesList(t *testing.T) {
	images := &Images{}
	results, err := images.List()
	if err != nil {
		t.Errorf("Images List failed with %s.", err)
	}
	if results == nil {
		t.Error("Images List results should not be nil.")
		return
	}
	for _, result := range results {
		if len(result.Id()) < 40 {
			t.Error("Wrong value for id.")
		}
		if TestImage.Name() != result.Name() {
			t.Errorf("Expected %s, but is %s", TestImage.Name(), result.Name())
		}
	}
}

func Test_ImagesShow(t *testing.T) {
	images := &Images{}
	result, err := images.Show(TestImage)
	if err != nil {
		t.Error("Images Show failed with %s.", err)
	}
	if result == nil {
		t.Error("Images Show result should not be nil.")
		return
	}
	if len(result.Id()) < 40 {
		t.Error("Wrong value for id.")
	}
	if TestImage.Name() != result.Name() {
		t.Errorf("Expected %s, but is %s", TestImage.Name(), result.Name())
	}
}

func Test_ImagesDestroy(t *testing.T) {
	images := &Images{}
	ok, err := images.Destroy(TestImage)
	if !ok {
		t.Error("Images Destroy failed.")
	}
	if err != nil {
		t.Error("Images Destroy failed with %s.", err)
	}
}

func Test_Done(t *testing.T) {
	if TestImageServer == nil {
		return
	}
	servers := &Servers{}
	servers.Destroy(TestImageServer)
}
