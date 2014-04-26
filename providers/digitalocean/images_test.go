// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package digitalocean

import (
	"fmt"
	"os"
	"testing"

	"github.com/gcloud/compute"
	"github.com/gcloud/identity"
)

var NewImage = &Image{data: &image{Name: "GCloudImage"}}
var TestImage compute.Image
var TestImageServer compute.Server

var iAccount = &identity.Account{
	Id:  "id",
	Key: "key",
}

func init() {
	provider = compute.GetProvider("digitalocean")
	provider.SetAccount(iAccount)
	r := []byte(`{
		"action_status": "done",
		"status": "OK",
		"droplet": {
			"id": 100824,
			"name": "GCloudServerOnDO",
			"image_id": 419,
			"size_id": 32,
			"event_id": 7499
		}
	}`)
	ts := newTestResponse(r)
	defer ts.Close()
	servers := &Servers{provider}
	result, err := servers.Create(servers.New(compute.Map{
		"name": NewImage.Name(),
	}))
	if result == nil {
		fmt.Printf("Create Server failed. %#v\n%#v\n", result, err)
		os.Exit(1)
	}
	TestImageServer = result
}

func NewImages() *Images {
	return &Images{provider}
}

func Test_ImagesCreate(t *testing.T) {
	r := []byte(`{
		"status": "OK",
		"event_id": 7504
	}`)
	ts := newTestResponse(r)
	defer ts.Close()
	images := NewImages()
	result, err := images.Create(images.New(compute.Map{"name": "GCloudImage"}))
	if err != nil {
		t.Errorf("Images Create failed with %s.", err)
	}
	fmt.Printf("%#v", result)
	if result == nil {
		t.Error("Images Create result should not be nil.")
		return
	}
	TestImage = result
	if result.Id() == "" {
		t.Error("Wrong value for id.")
	}
}

func Test_ImagesList(t *testing.T) {
	r := []byte(`{
		"status": "OK",
		"images": [
			{
				"id": 1,
				"name": "My first snapshot",
				"distribution": "Ubuntu",
				"slug": "ubuntu-12.10-x32",
				"public": true
			},
			{
				"id": 2,
				"name": "Automated Backup",
				"distribution": "Ubuntu"
			}
		]
	}`)
	ts := newTestResponse(r)
	defer ts.Close()
	images := NewImages()
	results, err := images.List()
	if err != nil {
		t.Errorf("Images List failed with %s.", err)
	}
	if results == nil {
		t.Error("Images List results should not be nil.")
		return
	}
	for _, result := range results {
		if result.Id() == "" {
			t.Error("Wrong value for id.")
		}
	}
}

func Test_ImagesShow(t *testing.T) {
	r := []byte(`{
		"status": "OK",
		"image": {
			"id": 2,
			"name": "Automated Backup",
			"distribution":"Ubuntu"
		}
	}`)
	ts := newTestResponse(r)
	defer ts.Close()
	images := NewImages()
	result, err := images.Show(NewImage)
	if err != nil {
		t.Errorf("Images Show failed with %s", err)
	}
	if result == nil {
		t.Error("Images Show result should not be nil.")
		return
	}
	if result.Id() == "" {
		t.Error("Wrong value for id.")
	}
}

func Test_ImagesDestroy(t *testing.T) {
	r := []byte(`{
 		"status": "OK"
	}`)
	ts := newTestResponse(r)
	defer ts.Close()
	images := NewImages()
	ok, err := images.Destroy(NewImage)
	if !ok {
		t.Error("Images Destroy failed.")
	}
	if err != nil {
		t.Errorf("Images Destroy failed with %s.", err)
	}
}
