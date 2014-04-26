// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	"testing"
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
func (m *MockImage) String() string {
	return ""
}
func (m *MockImage) MarshalJSON() ([]byte, error) {
	return []byte{}, nil
}

type MockImages struct{}

func (i *MockImages) New(m Map) Image {
	return &MockImage{}
}

func (i *MockImages) List() ([]Image, error) {
	results := make([]Image, 0)
	r := &MockImage{name: "My Image", id: "616fb98f-46ca-475e-917e-2563e5a8cd19"}
	return append(results, r), nil
}
func (i *MockImages) Show(Image) (Image, error) {
	r := &MockImage{name: "My Image", id: "616fb98f-46ca-475e-917e-2563e5a8cd19"}
	return r, nil
}
func (i *MockImages) Create(Image) (Image, error) {
	r := &MockImage{name: "My Image", id: "616fb98f-46ca-475e-917e-2563e5a8cd19"}
	return r, nil
}
func (i *MockImages) Destroy(Image) (bool, error) {
	return true, nil
}

func init() {
	RegisterImages("mock", &MockImages{})
}

func Test_ImagesList(t *testing.T) {
	images := GetImages("mock", nil)
	results, err := images.List()
	if err != nil {
		t.Error("Images List failed with %s.", err)
	}
	if results == nil {
		t.Error("Images List results should not be nil.")
		return
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
	images := GetImages("mock", nil)
	result, err := images.Show(&MockImage{id: "616fb98f-46ca-475e-917e-2563e5a8cd19"})
	if err != nil {
		t.Error("Images Show failed with %s.", err)
	}
	if result == nil {
		t.Error("Images Show results should not be nil.")
		return
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
	images := GetImages("mock", nil)
	result, err := images.Create(&MockImage{name: "My Image"})
	if err != nil {
		t.Error("Images Create failed with %s.", err)
	}
	if result == nil {
		t.Error("Images Create results should not be nil.")
		return
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
	images := GetImages("mock", nil)
	ok, err := images.Destroy(&MockImage{id: "616fb98f-46ca-475e-917e-2563e5a8cd19"})
	if !ok {
		t.Error("Images Destroy failed.")
	}
	if err != nil {
		t.Error("Images Destroy failed with " + err.Error() + ".")
	}
}
