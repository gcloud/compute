// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	"testing"

	p "github.com/gcloud/providers"
)

type MockLocation struct {
	id     string
	name   string
	region string
}

func (m *MockLocation) Id() string {
	return m.id
}
func (m *MockLocation) Name() string {
	return m.name
}
func (m *MockLocation) Region() string {
	return "east"
}

type MockLocations struct{}

// List locations available on the account.
func (l *MockLocations) List() ([]p.Location, error) {
	return []p.Location{
		&MockLocation{name: "The Location", id: "616fb98f-46ca-475e-917e-2563e5a8cd19"},
	}, nil
}
func (l *MockLocations) Show(id string) (p.Location, error) {
	return &MockLocation{name: "The Location", id: "616fb98f-46ca-475e-917e-2563e5a8cd19"}, nil
}

func init() {
	p.RegisterLocations("mock", &MockLocations{})
}

func Test_LocationsList(t *testing.T) {
	locations := &Locations{Provider: "mock"}
	results, err := locations.List()
	if err != nil {
		t.Error("Locations List failed with " + err.Error() + "(bool, error).")
	}
	if results == nil {
		t.Error("Results should not be nil.")
	}
	for _, v := range results {
		if v.Id() != "616fb98f-46ca-475e-917e-2563e5a8cd19" {
			t.Error("Wrong value for Id.")
		}
		if v.Name() != "The Location" {
			t.Error("Wrong value for Name.")
		}
	}
}

func Test_LocationsShow(t *testing.T) {
	locations := &Locations{Provider: "mock"}
	result, err := locations.Show("616fb98f-46ca-475e-917e-2563e5a8cd19")
	if err != nil {
		t.Error("Locations Show failed with " + err.Error() + ".")
	}
	if result == nil {
		t.Error("Results should not be nil.")
	}
	r := result
	if r.Id() != "616fb98f-46ca-475e-917e-2563e5a8cd19" {
		t.Error("Wrong value for Id.")
	}
	if r.Name() != "The Location" {
		t.Error("Wrong value for Name.")
	}
}
