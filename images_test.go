// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	p "github.com/gcloud/compute/providers"
	"testing"
)

type MockImages struct{}

// List images available on the account.
func (i *MockImages) List() ([]byte, error) {
	return []byte(`[{"Name":"My Image","Id": "616fb98f-46ca-475e-917e-2563e5a8cd19"}]`), nil
}
func (i *MockImages) Show(id string) ([]byte, error) {
	return []byte(`{"Name":"My Image","Id": "616fb98f-46ca-475e-917e-2563e5a8cd19"}`), nil
}
func (i *MockImages) Create(n *p.Image) ([]byte, error) {
	return []byte(`{"Name":"My Image","Id": "616fb98f-46ca-475e-917e-2563e5a8cd19"}`), nil
}
func (i *MockImages) Destroy(id string) (bool, error) {
	return true, nil
}

func init() {
	p.RegisterImages("mock", &MockImages{})
}

func Test_ImagesList(t *testing.T) {
	images := &Images{Provider: "mock"}
	results, err := images.List()
	if err != nil {
		t.Error("Images List failed with " + err.Error() + "(bool, error).")
	}
	if results == nil {
		t.Error("Results should not be nil.")
	}
	for _, v := range *results {
		if v.Id != "616fb98f-46ca-475e-917e-2563e5a8cd19" {
			t.Error("Wrong value for Id.")
		}
		if v.Name != "My Image" {
			t.Error("Wrong value for Name.")
		}
	}
}

func Test_ImagesShow(t *testing.T) {
	images := &Images{Provider: "mock"}
	result, err := images.Show("616fb98f-46ca-475e-917e-2563e5a8cd19")
	if err != nil {
		t.Error("Images Show failed with " + err.Error() + ".")
	}
	if result == nil {
		t.Error("Results should not be nil.")
	}
	r := *result
	if r.Id != "616fb98f-46ca-475e-917e-2563e5a8cd19" {
		t.Error("Wrong value for Id.")
	}
	if r.Name != "My Image" {
		t.Error("Wrong value for Name.")
	}
}

func Test_ImagesCreate(t *testing.T) {
	images := &Images{Provider: "mock"}
	result, err := images.Create(&p.Image{
		Name: "My Image"})
	if err != nil {
		t.Error("Images Create failed with " + err.Error() + ".")
	}
	if result == nil {
		t.Error("Results should not be nil.")
	}
	r := *result
	if r.Id != "616fb98f-46ca-475e-917e-2563e5a8cd19" {
		t.Error("Wrong value for Id.")
	}
	if r.Name != "My Image" {
		t.Error("Wrong value for Name.")
	}
}

func Test_ImagesDestroy(t *testing.T) {
	images := &Images{Provider: "mock"}
	ok, err := images.Destroy("616fb98f-46ca-475e-917e-2563e5a8cd19")
	if !ok {
		t.Error("Images Destroy failed.")
	}
	if err != nil {
		t.Error("Images Destroy failed with " + err.Error() + ".")
	}
}
