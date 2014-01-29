// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	"testing"

	p "github.com/gcloud/providers"
)

type MockImage struct {
	id   string
	name string
}

func (m *MockImage) Id() string {
	return m.id
}
func (m *MockImage) Name() string {
	return m.name
}
func (m *MockImage) File() string {
	return "file"
}

type MockImages struct{}

func (i *MockImages) List() ([]p.Image, error) {
	results := make([]p.Image, 0)
	r := &MockImage{name: "My Image", id: "616fb98f-46ca-475e-917e-2563e5a8cd19"}
	return append(results, r), nil
}
func (i *MockImages) Show(id string) (p.Image, error) {
	r := &MockImage{name: "My Image", id: "616fb98f-46ca-475e-917e-2563e5a8cd19"}
	return r, nil
}
func (i *MockImages) Create(n interface{}) (p.Image, error) {
	r := &MockImage{name: "My Image", id: "616fb98f-46ca-475e-917e-2563e5a8cd19"}
	return r, nil
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
	for _, v := range results {
		if v.Id() != "616fb98f-46ca-475e-917e-2563e5a8cd19" {
			t.Error("Wrong value for Id.")
		}
		if v.Name() != "My Image" {
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
	r := result
	if r.Id() != "616fb98f-46ca-475e-917e-2563e5a8cd19" {
		t.Error("Wrong value for Id.")
	}
	if r.Name() != "My Image" {
		t.Error("Wrong value for Name.")
	}
}

func Test_ImagesCreate(t *testing.T) {
	images := &Images{Provider: "mock"}
	result, err := images.Create(MockImage{name: "My Image"})
	if err != nil {
		t.Error("Images Create failed with " + err.Error() + ".")
	}
	if result == nil {
		t.Error("Results should not be nil.")
	}
	r := result
	if r.Id() != "616fb98f-46ca-475e-917e-2563e5a8cd19" {
		t.Error("Wrong value for Id.")
	}
	if r.Name() != "My Image" {
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
